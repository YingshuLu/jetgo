package native

import (
	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value"
	"github.com/yingshulu/jetgo/value/array"
)

func inArrayContains(x interface{}, arr []interface{}) bool {
	for _, v := range arr {
		if x == v {
			return true
		}
		if value.Equal(x, v) {
			return true
		}
	}
	return false
}

func inArray(args []interface{}) (interface{}, error) {
	return args, nil
}

func inArrayAppend(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, consts.ErrArgsNotEnough(args)
	}
	return array.Append(args[0], args[1])
}
