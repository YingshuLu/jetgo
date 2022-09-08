// Package number value of expression
package number

import (
	"fmt"
	"go/token"
	"math"

	"github.com/yingshulu/jetgo/consts"
)

const (
	epsNumber   = 0.00000001
	emptyNumber = 0.0
)

// operator number basic operator method
type operator func(l, r float64) interface{}

var (
	operatorRoute = []operator{
		token.ADD: add,
		token.SUB: sub,
		token.MUL: mul,
		token.QUO: div,
		token.REM: mod,

		// bitwise
		token.AND: and,
		token.OR:  or,
		token.XOR: xor,
		token.SHL: leftShift,
		token.SHR: rightShift,

		// comparison
		token.EQL: equal,
		token.LSS: less,
		token.GTR: greater,
		token.NEQ: notEqual,
		token.LEQ: lessEqual,
		token.GEQ: greaterEqual,
	}
	routeEnd = len(operatorRoute)
)

// Type check if any is Number
func Type(any interface{}) bool {
	switch any.(type) {
	case float32, float64:
		return true
	case int8, int16, int32, int64, int:
		return true
	case uint8, uint16, uint32, uint64, uint:
		return true
	case *float32, *float64:
		return true
	case *int8, *int16, *int32, *int64, *int:
		return true
	case *uint8, *uint16, *uint32, *uint64, *uint:
		return true
	default:
		return false
	}
}

// Eval evaluate value of right and left with operator method
func Eval(left, right interface{}, op token.Token) (interface{}, error) {
	l, err := Number(left)
	if err != nil {
		return nil, err
	}
	r, err := Number(right)
	if err != nil {
		return nil, err
	}
	err = checkException(l, r, op)
	if err != nil {
		return nil, err
	}
	if int(op) >= routeEnd {
		return nil, consts.ErrUndefinedType(op)
	}
	f := operatorRoute[op]
	if f != nil {
		return f(l, r), nil
	}
	return nil, consts.ErrUndefinedToken(op)
}

func checkException(l, r float64, op token.Token) error {
	if (op == token.QUO || op == token.REM) && equalNumber(r, 0) {
		return consts.ErrMathUndefined(fmt.Sprintf("%v %v %v", l, op, r))
	}
	return nil
}

// arithmetic
func add(l, r float64) interface{} {
	return l + r
}

func sub(l, r float64) interface{} {
	return l - r
}

func mul(l, r float64) interface{} {
	return l * r
}

func div(l, r float64) interface{} {
	return l / r
}

func mod(l, r float64) interface{} {
	return math.Mod(l, r)
}

// bitwise
func and(l, r float64) interface{} {
	return float64(int64(l) & int64(r))
}

func or(l, r float64) interface{} {
	return float64(int64(l) | int64(r))
}

func xor(l, r float64) interface{} {
	return float64(int64(l) ^ int64(r))
}

func leftShift(l, r float64) interface{} {
	return float64(int64(l) << int64(r))
}

func rightShift(l, r float64) interface{} {
	return float64(int64(l) >> int64(r))
}

// comparison
func equal(l, r float64) interface{} {
	return equalNumber(l, r)
}

func less(l, r float64) interface{} {
	return lessNumber(l, r)
}

func greater(l, r float64) interface{} {
	return greaterNumber(l, r)
}

func notEqual(l, r float64) interface{} {
	return !equalNumber(l, r)
}

func lessEqual(l, r float64) interface{} {
	return lessNumber(l, r) || equalNumber(l, r)
}

func greaterEqual(l, r float64) interface{} {
	return greaterNumber(l, r) || equalNumber(l, r)
}

func equalNumber(l, r float64) bool {
	return math.Abs(l-r) < epsNumber
}

func lessNumber(l, r float64) bool {
	return l-r < -epsNumber
}

func greaterNumber(l, r float64) bool {
	return l-r > epsNumber
}
