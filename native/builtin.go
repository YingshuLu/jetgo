package native

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value"
	"github.com/yingshulu/jetgo/value/array"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
	"github.com/yingshulu/jetgo/value/table"
)

func inLen(args []interface{}) (interface{}, error) {
	if len(args) == 0 {
		return nil, consts.ErrArgsNotEnough(fmt.Sprintf("func %v, args %v", "length", 0))
	}

	// string type
	ok := stringv.Type(args[0])
	if ok {
		return stringv.Len(args[0])
	}

	// table type
	t, ok := args[0].(map[string]interface{})
	if ok {
		return len(t), nil
	}

	// array type
	return array.Len(args[0])
}

func inMethod(args []interface{}) (interface{}, error) {
	xt := value.Type(args[0])
	if xt == value.UnknownType {
		return false, fmt.Errorf("func in args[0] not builtin type: %T", args[0])
	}
	switch arr := args[1].(type) {
	case []interface{}:
		return inArrayContains(args[0], arr), nil
	case map[string]interface{}:
		return table.In(args[0], args[1])
	}
	if array.Type(args[1]) {
		return array.In(args[0], args[1])
	}
	return false, fmt.Errorf("func in args[1] not array type: %T", args[0])
}

func toNumber(args []interface{}) (interface{}, error) {
	if len(args) < 1 {
		return nil, consts.ErrArgsNotEnough(args)
	}

	any := args[0]
	v, ok := number.Value(any)
	if ok {
		return v, nil
	}

	s, err := stringv.String(any)
	if err != nil {
		return nil, err
	}
	return strconv.ParseFloat(s, 64)
}

func toInteger(args []interface{}) (interface{}, error) {
	v, err := toNumber(args)
	if err != nil {
		return nil, err
	}

	switch fv := v.(type) {
	case float32:
		return int64(fv), nil
	case float64:
		return int64(fv), nil
	}
	return v, nil
}

func inTime(_ []interface{}) (interface{}, error) {
	return time.Now().Unix(), nil
}

func inLog(args []interface{}) (interface{}, error) {
	log.Print(args)
	return nil, nil
}
