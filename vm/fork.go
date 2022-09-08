package vm

import (
	"fmt"
	"go/ast"
	"go/token"
	"sync"
	"time"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value/number"
)

var vmPool = &sync.Pool{
	New: func() interface{} {
		return &vm{
			callArgsStack: make(stack, 8),
		}
	},
}

func fetchVM() *vm {
	return vmPool.Get().(*vm)
}

func backVM(v *vm) {
	v.symbolStack = nil
	v.rule = nil
	v.err = nil
	v.branch = token.ILLEGAL
	vmPool.Put(v)
}

func (v *vm) forkLen(typ *ast.ArrayType, fact interface{}) (float64, error) {
	ident, ok := typ.Elt.(*ast.Ident)
	if !ok {
		return 0, consts.ErrSemanticUndefined(ident.Pos())
	}
	if ident.Name != forkToken {
		return 0, consts.ErrSemanticUndefined(ident.Pos())
	}

	val, err := v.evalExpr(typ.Len, fact)
	if err != nil {
		return 0, err
	}
	fv, err := number.Number(val)
	if !ok {
		return 0, err
	}
	return fv, nil
}

func (v *vm) forkExpr(compose *ast.CompositeLit, fact interface{}) (interface{}, error) {
	arr, ok := compose.Type.(*ast.ArrayType)
	if !ok {
		return nil, consts.ErrSemanticUndefined(compose.Pos())
	}

	to, err := v.forkLen(arr, fact)
	if err != nil {
		return nil, err
	}
	timeout := time.Duration(to) * time.Millisecond

	n := len(compose.Elts)
	results := make([]interface{}, n)
	ch := make(chan forkResult, n)

	for index, expr := range compose.Elts {
		// race: clone vm context
		nv := v.clone()
		go forkGo(nv, expr, fact, ch, index)
	}

	// race: merge results
	var beTimeout bool
	for i := 0; i < n && !beTimeout; i++ {
		select {
		case re := <-ch:
			fe, ok := re.val.(error)
			if ok {
				return nil, fe
			}
			results[re.index] = re.val
		case <-time.After(timeout):
			beTimeout = true
		}
	}

	// race: detach shares due to timeout happened
	if beTimeout {
		v.detach()
	}
	return results, nil
}

func (v *vm) clone() *vm {
	nv := fetchVM()

	// share: code segment
	nv.rule = v.rule
	nv.branch = v.branch
	nv.err = v.err

	// share: data segment
	nv.symbolStack = v.symbolStack

	// share: method call stack
	n := nv.rule.CallArgs()
	nv.callArgsStack = nv.callArgsStack.growNoCopy(n)
	return nv
}

func (v *vm) detach() {
	// copy a new data segment
	s := make(stack, len(v.symbolStack))
	copy(s, v.symbolStack)
	v.symbolStack = s
}

func (v *vm) debugInfo() string {
	return fmt.Sprintf(`
	+ rule name: %s
		- consts    : %v
		- vars count: %d
		- args count: %d
		- source    : %s

	+ vm addr: %p
		- vars stack: %v
		- args stack: %v
			`,
		v.rule.Name(), v.rule.Consts(),
		len(v.rule.SymbolTable()), v.rule.CallArgs(), v.rule.Source(),
		v, v.symbolStack, v.callArgsStack)
}

type forkResult struct {
	index int
	val   interface{}
}

func forkGo(v *vm, fn ast.Expr, fact interface{}, ch chan forkResult, index int) {
	defer backVM(v)

	r := forkResult{index: index}
	val, err := v.evalExpr(fn, fact)
	if err != nil {
		r.val = err
	} else {
		r.val = val
	}
	ch <- r
}
