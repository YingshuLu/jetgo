package native

import (
	"fmt"
	"go/token"
	"strconv"
	"strings"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/number"

	stringv "github.com/yingshulu/jetgo/value/string"
)

func toString(args []interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", nil
	}

	any := args[0]
	s, ok := stringv.Value(any)
	if ok {
		return s, nil
	}

	v, ok := number.Value(any)
	if ok {
		return fmt.Sprintf("%v", v), nil
	}

	b, ok := boolean.Value(any)
	if ok {
		return strconv.FormatBool(b), nil
	}

	return fmt.Sprintf("%v", any), nil
}

func matchString(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return false, nil
	}
	res, err := stringv.Eval(args[0], args[1], token.SHR)
	if err != nil {
		return false, err
	}
	return res, nil
}

func containString(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return false, nil
	}
	left, err := stringv.String(args[0])
	if err != nil {
		return false, err
	}
	right, err := stringv.String(args[1])
	if err != nil {
		return false, err
	}
	return strings.Contains(left, right), nil
}

func toSupper(args []interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", nil
	}

	any := args[0]
	s, ok := stringv.Value(any)
	if !ok {
		return nil, fmt.Errorf("supper not string but %t", any)
	}
	return strings.ToUpper(s), nil
}

func toSlower(args []interface{}) (interface{}, error) {
	if len(args) < 1 {
		return "", nil
	}

	any := args[0]
	s, ok := stringv.Value(any)
	if !ok {
		return nil, fmt.Errorf("supper not string but %t", any)
	}
	return strings.ToLower(s), nil
}

func inSplit(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, consts.ErrArgsNotEnough(args)
	}

	s, ok := stringv.Value(args[0])
	if !ok {
		return nil, fmt.Errorf("args[0] not string but %t", args[0])
	}
	sp, ok := stringv.Value(args[1])
	if !ok {
		return nil, fmt.Errorf("args[1] not string but %t", args[1])
	}
	return strings.Split(s, sp), nil
}

func inSprint(args []interface{}) (interface{}, error) {
	if len(args) < 1 {
		return nil, nil
	}
	return fmt.Sprint(args...), nil
}

func inSprintf(args []interface{}) (interface{}, error) {
	n := len(args)
	if n < 1 {
		return nil, nil
	}
	fmtStr, ok := args[0].(string)
	if !ok || n == 1 {
		return inSprint(args)
	}
	return fmt.Sprintf(fmtStr, args[1:]...), nil
}
