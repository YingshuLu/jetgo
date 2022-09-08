// Package boolean bool value of expression
package boolean

import (
	"go/token"

	"github.com/yingshulu/jetgo/consts"
)

// operator boolean basic operator method
type operator func(l, r bool) interface{}

var (
	operatorRoute = []operator{
		token.LAND: and,
		token.LOR:  or,
		token.EQL:  equal,
		token.NEQ:  notEqual,
	}
	routeEnd = len(operatorRoute)
)

// Type check if any is boolean
func Type(any interface{}) bool {
	switch any.(type) {
	case bool:
		return true
	case *bool:
		return true
	default:
		return false
	}
}

// Boolean convert any into boolean if success
func Boolean(any interface{}) (bool, bool) {
	return Value(any)
}

// Value check if any is bool, return value if it is
func Value(any interface{}) (bool, bool) {
	switch v := any.(type) {
	case bool:
		return v, true
	case *bool:
		if v == nil {
			return false, true
		}
		return *v, true
	default:
		return false, false
	}
}

// Eval result of boolean expression
func Eval(left, right interface{}, op token.Token) (interface{}, error) {
	l, ok := Boolean(left)
	if !ok {
		return nil, consts.ErrNotBool(left)
	}
	r, ok := Boolean(right)
	if !ok {
		return nil, consts.ErrNotBool(right)
	}

	if int(op) >= routeEnd {
		return nil, consts.ErrUndefinedToken(op)
	}
	f := operatorRoute[op]
	if f != nil {
		return f(l, r), nil
	}
	return nil, consts.ErrUndefinedToken(op)
}

func and(l, r bool) interface{} {
	return l && r
}

func or(l, r bool) interface{} {
	return l || r
}

func equal(l, r bool) interface{} {
	return l == r
}

func notEqual(l, r bool) interface{} {
	return l != r
}
