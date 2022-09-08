package vm

import (
	"go/ast"
	"go/token"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value/boolean"
)

func (v *vm) forStmt(stmt *ast.ForStmt, fact interface{}) (interface{}, error) {
	var (
		loop = true
		err  error
		val  interface{}
		ok   bool
	)

	// check init
	if stmt.Init != nil {
		_, err = v.evalStmt(stmt.Init, fact)
	}
	if err != nil {
		return nil, err
	}

	for {
		// check condition
		if stmt.Cond != nil {
			val, err = v.evalExpr(stmt.Cond, fact)
			if err != nil {
				return nil, err
			}
			loop, ok = boolean.Value(val)
			if !ok {
				return nil, consts.ErrNotBool(val)
			}
		}
		if !loop {
			break
		}

		// check body
		val, err = v.evalStmt(stmt.Body, fact)
		if err != nil {
			return nil, err
		}
		// body returned
		if val != nil {
			return val, nil
		}

		// check break
		if v.branch == token.BREAK {
			loop = false
		}
		v.branch = token.ILLEGAL
		if !loop {
			break
		}

		// check post
		if stmt.Post != nil {
			_, err = v.evalStmt(stmt.Post, fact)
		}
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}
