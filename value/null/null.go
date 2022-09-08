// Package null value of expression
package null

import (
	"go/token"
	"reflect"

	"github.com/yingshulu/jetgo/consts"
)

// Type check if any is bool type
func Type(any interface{}) bool {
	return IsNil(any)
}

// Eval of left and right with token
func Eval(l, r interface{}, op token.Token) (interface{}, error) {
	switch op {
	case token.EQL:
		return IsNil(l) == IsNil(r), nil
	case token.NEQ:
		return IsNil(l) != IsNil(r), nil
	default:
		return false, consts.ErrUndefinedToken(op)
	}
}

// IsNil check if any is nil
func IsNil(any interface{}) bool {
	if any == nil {
		return true
	}
	switch reflect.TypeOf(any).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Slice:
		return reflect.ValueOf(any).IsNil()
	default:
		return false
	}
}
