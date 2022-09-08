package vm

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/native"
	"github.com/yingshulu/jetgo/value/array"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/null"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
	"github.com/yingshulu/jetgo/value/table"
)

func (v *vm) evalExpr(expr ast.Expr, fact interface{}) (interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}

	var (
		res interface{}
		err error
	)
	switch node := expr.(type) {
	case *ast.BinaryExpr:
		res, err = v.binaryExpr(node, fact)
	case *ast.BasicLit:
		res, err = v.basicLit(node, fact)
	case *ast.Ident:
		res, err = v.ident(node, fact)
	case *ast.UnaryExpr:
		res, err = v.unaryExpr(node, fact)
	case *ast.ParenExpr:
		res, err = v.parenExpr(node, fact)
	case *ast.SelectorExpr:
		res, err = v.selectorExpr(node, fact)
	case *ast.IndexExpr:
		res, err = v.indexExpr(node, fact)
	case *ast.CallExpr:
		res, err = v.callExpr(node, fact)
	case *ast.FuncLit:
		res, err = v.funcLit(node, fact)
	case *ast.CompositeLit:
		res, err = v.composeExpr(node, fact)
	default:
		err = consts.ErrSemanticUndefined(node)
	}

	// deep visit first
	if err != nil && v.err == nil {
		file := v.rule.File()
		pos := (*file).Position(expr.Pos())
		v.err = fmt.Errorf("[%s] %v", pos.String(), err)
	}
	return res, v.err
}

func (v *vm) evalStmt(node ast.Stmt, fact interface{}) (interface{}, error) {
	if v.err != nil {
		return nil, v.err
	}

	var (
		res interface{}
		err error
	)
	switch nd := node.(type) {
	case *ast.ReturnStmt:
		res, err = v.returnStmt(nd, fact)
	case *ast.AssignStmt:
		res, err = v.assignStmt(nd, fact)
	case *ast.IfStmt:
		res, err = v.ifStmt(nd, fact)
	case *ast.BlockStmt:
		res, err = v.blockStmt(nd, fact)
	case *ast.ExprStmt:
		res, err = v.exprStmt(nd, fact)
	case *ast.ForStmt:
		res, err = v.forStmt(nd, fact)
	case *ast.BranchStmt:
		res, err = v.branchStmt(nd, fact)
	case nil:
	default:
		err = consts.ErrSemanticUndefined(node)
	}

	// deep visit first
	if err != nil && v.err == nil {
		file := v.rule.File()
		pos := (*file).Position(node.Pos())
		v.err = fmt.Errorf("[%s] %v", pos.String(), err)
	}
	return res, v.err
}

func (v *vm) fetch(key string, data interface{}) (interface{}, error) {
	if varRenamed(key) {
		idx := indexing(key)
		return v.symbolStack[idx], nil
	}
	fact, ok := data.(map[string]interface{})
	if !ok {
		return nil, consts.ErrFactNotFound(key)
	}

	val, ok := fact[key]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func skipBinary(left interface{}, op token.Token) bool {
	if op != token.LOR && op != token.LAND {
		return false
	}
	b, ok := boolean.Boolean(left)
	if !ok {
		return false
	}
	if op == token.LOR && b {
		return true
	}
	if op == token.LAND && !b {
		return true
	}
	return false
}

func (v *vm) binaryExpr(node *ast.BinaryExpr, fact interface{}) (interface{}, error) {
	xn, yn := node.X, node.Y

	// ident first evaluated for acceleration
	if logicToken(node.Op) {
		if _, ok := yn.(*ast.Ident); ok {
			xn, yn = yn, xn
		}
	}

	left, err := v.evalExpr(xn, fact)
	if err != nil {
		return nil, err
	}

	if skipBinary(left, node.Op) {
		return left, nil
	}

	right, err := v.evalExpr(yn, fact)
	if err != nil {
		return nil, err
	}

	if null.Type(right) || null.Type(left) {
		return null.Eval(left, right, node.Op)
	}
	if number.Type(left) && number.Type(right) {
		return number.Eval(left, right, node.Op)
	}
	if stringv.Type(left) && stringv.Type(right) {
		return stringv.Eval(left, right, node.Op)
	}
	if boolean.Type(left) && boolean.Type(right) {
		return boolean.Eval(left, right, node.Op)
	}

	return nil, consts.ErrUndefinedType(fmt.Sprintf("%v %v %v", left, node.Op, right))
}

func (v *vm) basicLit(node *ast.BasicLit, _ interface{}) (interface{}, error) {
	return v.rule.Consts()[indexing(node.Value)], nil
}

func (v *vm) ident(node *ast.Ident, fact interface{}) (interface{}, error) {
	switch node.Name {
	case trueToken:
		return true, nil
	case falseToken:
		return false, nil
	case nullToken:
		return nil, nil
	default:
		return v.fetch(node.Name, fact)
	}
}

// only support ! - token
func (v *vm) unaryExpr(node *ast.UnaryExpr, fact interface{}) (interface{}, error) {
	x, err := v.evalExpr(node.X, fact)
	if err != nil {
		return nil, err
	}
	switch node.Op {
	case token.NOT:
		val, ok := boolean.Boolean(x)
		if !ok {
			return nil, consts.ErrNotBool(x)
		}
		return !val, nil
	case token.SUB:
		val, err := number.Number(x)
		if err != nil {
			return nil, err
		}
		return -1 * val, nil
	default:
		return nil, consts.ErrUndefinedToken(node.Op)
	}
}

func (v *vm) parenExpr(node *ast.ParenExpr, fact interface{}) (interface{}, error) {
	return v.evalExpr(node.X, fact)
}

func (v *vm) selectorExpr(node *ast.SelectorExpr, fact interface{}) (interface{}, error) {
	base, err := v.evalExpr(node.X, fact)
	if err != nil {
		return nil, err
	}
	return v.fetch(node.Sel.Name, base)
}

func (v *vm) indexExpr(node *ast.IndexExpr, fact interface{}) (interface{}, error) {
	val, err := v.evalExpr(node.X, fact)
	if err != nil {
		return nil, err
	}
	idx, err := v.evalExpr(node.Index, fact)
	if err != nil {
		return nil, err
	}

	// table
	if table.Type(val) {
		return table.Get(val, idx)
	}

	// array
	n, err := number.Number(idx)
	if err != nil {
		return nil, err
	}

	return array.Get(val, int(n))
}

func (v *vm) callExpr(node *ast.CallExpr, fact interface{}) (interface{}, error) {
	name := node.Fun.(*ast.Ident).Name
	method, _ := native.Fetch(name)
	end := len(node.Args)
	if end == 0 {
		return method(nil)
	}
	var (
		args stack
		val  interface{}
		err  error
	)

	// recover call stack
	args = v.callArgsStack
	defer func(stk stack) {
		v.callArgsStack = stk
	}(args)

	// allocate call stack
	callArgs := args[0:end]
	v.callArgsStack = args[end:]

	// exception for array method
	if name == arrayToken {
		size := end
		if size == 0 {
			size = 2
		}
		callArgs = make(stack, size)
	}

	for idx, nd := range node.Args {
		val, err = v.evalExpr(nd, fact)
		if err != nil {
			return nil, err
		}
		callArgs[idx] = val
	}
	return method(callArgs)
}

func (v *vm) composeExpr(node *ast.CompositeLit, fact interface{}) (interface{}, error) {
	switch node.Type.(type) {
	case *ast.ArrayType:
		return v.forkExpr(node, fact)
	case *ast.Ident:
		return v.tableExpr(node, fact)
	default:
		return nil, consts.ErrSemanticUndefined(node.Pos())
	}
}

func (v *vm) funcLit(node *ast.FuncLit, fact interface{}) (interface{}, error) {
	return v.blockStmt(node.Body, fact)
}

// stmt node travel //
func (v *vm) blockStmt(node *ast.BlockStmt, fact interface{}) (interface{}, error) {
	var (
		val interface{}
		err error
	)
	for _, stmt := range node.List {
		val, err = v.evalStmt(stmt, fact)

		// error or return happened
		if err != nil || val != nil {
			return val, err
		}

		// for block stop
		if v.branch == token.BREAK || v.branch == token.CONTINUE {
			return nil, nil
		}
	}
	return nil, nil
}

func (v *vm) assignStmt(node *ast.AssignStmt, fact interface{}) (interface{}, error) {
	val, err := v.evalExpr(node.Rhs[0], fact)
	if err != nil {
		return nil, err
	}
	lnd, ok := node.Lhs[0].(*ast.Ident)
	if !ok {
		return nil, consts.ErrSemanticUndefined(node.TokPos)
	}
	idx := indexing(lnd.Name)
	v.symbolStack[idx] = val
	return nil, nil
}

func (v *vm) returnStmt(node *ast.ReturnStmt, fact interface{}) (interface{}, error) {
	if len(node.Results) == 0 {
		return nil, nil
	}
	return v.evalExpr(node.Results[0], fact)
}

func (v *vm) ifStmt(node *ast.IfStmt, fact interface{}) (interface{}, error) {
	val, err := v.evalExpr(node.Cond, fact)
	if err != nil {
		return nil, err
	}
	cond, ok := boolean.Boolean(val)
	if !ok {
		return nil, consts.ErrNotBool(val)
	}
	if cond {
		return v.evalStmt(node.Body, fact)
	}
	return v.evalStmt(node.Else, fact)
}

// not care returned value
func (v *vm) exprStmt(node *ast.ExprStmt, fact interface{}) (interface{}, error) {
	callNode, ok := node.X.(*ast.CallExpr)
	if !ok {
		return nil, consts.ErrSemanticUndefined(node.X.Pos())
	}
	_, err := v.evalExpr(callNode, fact)
	return nil, err
}

func (v *vm) branchStmt(node *ast.BranchStmt, fact interface{}) (interface{}, error) {
	v.branch = node.Tok
	return nil, nil
}
