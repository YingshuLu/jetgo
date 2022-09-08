package consts

import (
	"errors"
	"fmt"
	"go/token"
)

// type error
var (
	// ErrNotBool 非Bool值类型错误
	ErrNotBool = func(v interface{}) error {
		return fmt.Errorf("not bool value: %v", v)
	}

	// ErrNotNumber 非Number值类型错误
	ErrNotNumber = func(v interface{}) error {
		return fmt.Errorf("not Number value: %v", v)
	}

	// ErrNotString 非String值类型错误
	ErrNotString = func(v interface{}) error {
		return fmt.Errorf("not string value: %v", v)
	}

	// ErrNotSlice 非Array值类型错误
	ErrNotSlice = func(v interface{}) error {
		return fmt.Errorf("not Array value: %v", v)
	}

	// ErrNotTable 非Table值类型错误
	ErrNotTable = func(v interface{}) error {
		return fmt.Errorf("not Table value: %v", v)
	}

	// ErrMethodNotFound 未定义Method错误
	ErrMethodNotFound = func(v interface{}) error {
		return fmt.Errorf("method not found: %v", v)
	}

	// ErrUndefinedType 未定义数据类型错误
	ErrUndefinedType = func(v interface{}) error {
		return fmt.Errorf("unsupport type: %v", v)
	}

	// ErrMathUndefined 未定义运算符错误
	ErrMathUndefined = func(v interface{}) error {
		return fmt.Errorf("math undefined: %v", v)
	}

	// ErrNotArithToken 未定义Arith符错误
	ErrNotArithToken = func(v token.Token) error {
		return fmt.Errorf("not arithmetic token: %v", v)
	}

	// ErrorNotCompToken 未定义比较符错误
	ErrorNotCompToken = func(v token.Token) error {
		return fmt.Errorf("not comparison token: %v", v)
	}

	// ErrUndefinedToken 未定义token错误
	ErrUndefinedToken = func(v token.Token) error {
		return fmt.Errorf("unsupport token: %v", v)
	}

	// ErrArgsNotEnough method 参数不足错误
	ErrArgsNotEnough = func(v interface{}) error {
		return fmt.Errorf("args not enough: %v", v)
	}

	// ErrFactNotFound Fact 数据未找到索引错误
	ErrFactNotFound = func(v interface{}) error {
		return fmt.Errorf("fact value: %v not found", v)
	}

	// ErrStackNotFound 临时栈值未找到错误
	ErrStackNotFound = func(v interface{}) error {
		return fmt.Errorf("stack value: %v not found", v)
	}

	// ErrSemanticUndefined 规则语法错误
	ErrSemanticUndefined = func(v interface{}) error {
		return fmt.Errorf("semantic undefined: %v", v)
	}
)

// compile error
var (
	// ErrVarNotFound 变量未找到错误
	ErrVarNotFound = func(v interface{}) error {
		return fmt.Errorf("variant not found: %v", v)
	}

	// ErrFuncNotFound 函数未找到错误
	ErrFuncNotFound = func(v interface{}) error {
		return fmt.Errorf("func not found: %v", v)
	}

	// ErrToManyLocalVars 临时变量数溢出错误
	ErrToManyLocalVars = errors.New("too many local variants")
)
