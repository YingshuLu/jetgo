// Package string value of expression
package string

import (
	"fmt"
	"go/token"
	"regexp"
	"sync"
	"unsafe"

	"github.com/yingshulu/jetgo/consts"
)

var (
	emptyString = ""
	regexpMap   sync.Map
)

// operator string basic operator method
type operator func(l, r string) (interface{}, error)

var (
	operatorRoute = []operator{
		token.ADD: add,      // +
		token.EQL: equal,    // ==
		token.NEQ: notEqual, // !=
		token.SHR: match,    // >>
	}
	routeEnd = len(operatorRoute)
)

// Type check if any is string type
func Type(any interface{}) bool {
	switch any.(type) {
	case string:
		return true
	case *string:
		return true
	case []byte:
		return true
	default:
		return false
	}
}

// String convert any to string if success
func String(any interface{}) (string, error) {
	v, ok := Value(any)
	if !ok {
		return emptyString, consts.ErrNotString(any)
	}
	return v, nil
}

// Value check if any is string
func Value(any interface{}) (string, bool) {
	switch v := any.(type) {
	case string:
		return v, true
	case *string:
		if v == nil {
			return emptyString, true
		}
		return *v, true
	case []byte:
		return *(*string)(unsafe.Pointer(&v)), true
	default:
		return emptyString, false
	}
}

// Eval string expression
func Eval(left, right interface{}, op token.Token) (interface{}, error) {
	l, err := String(left)
	if err != nil {
		return nil, err
	}
	r, err := String(right)
	if err != nil {
		return nil, consts.ErrNotString(r)
	}
	if int(op) >= routeEnd {
		return nil, consts.ErrUndefinedType(op)
	}
	f := operatorRoute[op]
	if f != nil {
		return f(l, r)
	}
	return nil, consts.ErrUndefinedToken(op)
}

// Len get string length
func Len(v interface{}) (int, error) {
	s, err := String(v)
	if err != nil {
		return 0, err
	}
	return len(s), nil
}

func add(l, r string) (interface{}, error) {
	return fmt.Sprintf("%s%s", l, r), nil
}

func equal(l, r string) (interface{}, error) {
	return l == r, nil
}

func notEqual(l, r string) (interface{}, error) {
	return l != r, nil
}

func match(target, pattern string) (interface{}, error) {
	re, err := getRegexp(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(target), nil
}

func getRegexp(pattern string) (*regexp.Regexp, error) {
	re := loadRegexpCache(pattern)
	if re != nil {
		return re, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	regexpMap.Store(pattern, re)
	return re, nil
}

func loadRegexpCache(pattern string) *regexp.Regexp {
	r, ok := regexpMap.Load(pattern)
	if ok {
		return r.(*regexp.Regexp)
	}
	return nil
}
