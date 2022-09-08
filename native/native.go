package native

import (
	"strings"

	"github.com/yingshulu/jetgo/value/method"
)

var methodMap = map[string]method.Method{
	"array":    inArray,
	"len":      inLen,
	"in":       inMethod,
	"append":   inArrayAppend,
	"string":   toString,
	"supper":   toSupper,
	"slower":   toSlower,
	"printf":   inSprintf,
	"sprintf":  inSprintf,
	"ssplit":   inSplit,
	"match":    matchString,
	"contains": containString,
	"number":   toNumber,
	"integer":  toInteger,
	"table":    inTable,
	"tset":     inTableSet,
	"tseti":    inTableSetInt,
	"tget":     inTableGet,
	"tdel":     inTableDel,
	"time":     inTime,
	"log":      inLog,
}

// Register method with name in express, NOTE: name should be lower case
func Register(name string, method method.Method) {
	methodMap[strings.ToLower(name)] = method
}

// Fetch method with name
func Fetch(name string) (method.Method, bool) {
	method, ok := methodMap[name]
	return method, ok
}
