package vm

import (
	"fmt"
	"go/ast"

	"github.com/yingshulu/jetgo/consts"
	stringv "github.com/yingshulu/jetgo/value/string"
)

func (v *vm) tableExpr(compose *ast.CompositeLit, fact interface{}) (interface{}, error) {
	indent, ok := compose.Type.(*ast.Ident)
	if !ok {
		return nil, consts.ErrSemanticUndefined(compose.Pos())
	}

	if indent.Name != tableToken {
		return nil, consts.ErrSemanticUndefined(compose.Pos())
	}

	var (
		table    = make(map[string]interface{}, len(compose.Elts))
		kv       *ast.KeyValueExpr
		err      error
		key, val interface{}
		sk       string
	)

	for _, node := range compose.Elts {
		kv, ok = node.(*ast.KeyValueExpr)
		if !ok {
			err = consts.ErrSemanticUndefined(node.Pos())
			break
		}
		key, err = v.evalExpr(kv.Key, fact)
		if err != nil {
			break
		}
		val, err = v.evalExpr(kv.Value, fact)
		if err != nil {
			break
		}
		sk, err = stringv.String(key)
		if err != nil {
			err = consts.ErrSemanticUndefined(fmt.Sprintf("pos: %v, message: %v", kv.Key.Pos(), err))
			break
		}
		table[sk] = val
	}
	return table, err
}
