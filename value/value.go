// Package value basic type values and expression
package value

import (
	"go/token"

	"github.com/yingshulu/jetgo/value/array"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/null"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
	"github.com/yingshulu/jetgo/value/table"
)

const (
	// UnknownType unknown value type
	UnknownType uint8 = iota

	// NullType null value type
	NullType

	// BooleanType boolean value type
	BooleanType

	// NumberType number value type
	NumberType

	// StringType number value type
	StringType

	// ArrayType array value type
	ArrayType

	// TableType table value type
	TableType
)

// Type of internal
func Type(v interface{}) uint8 {
	if number.Type(v) {
		return NumberType
	}
	if boolean.Type(v) {
		return BooleanType
	}
	if stringv.Type(v) {
		return StringType
	}
	if array.Type(v) {
		return ArrayType
	}
	if null.Type(v) {
		return NullType
	}
	if table.Type(v) {
		return TableType
	}
	return UnknownType
}

// Equal compare left with right
func Equal(left, right interface{}) bool {
	if left == right {
		return true
	}

	lt, rt := Type(left), Type(right)
	if lt != rt {
		return false
	}

	var (
		res interface{}
		err error
	)
	switch lt {
	case NumberType:
		res, err = number.Eval(left, right, token.EQL)
		if err != nil {
			return false
		}
		return res.(bool)
	case StringType:
		res, err = stringv.Eval(left, right, token.EQL)
		if err != nil {
			return false
		}
		return res.(bool)
	case BooleanType:
		res, err = boolean.Eval(left, right, token.EQL)
		if err != nil {
			return false
		}
		return res.(bool)
	default:
		return false
	}
}
