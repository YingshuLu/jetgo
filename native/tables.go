package native

import (
	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/number"
	"github.com/yingshulu/jetgo/value/table"
)

func inTable(args []interface{}) (interface{}, error) {
	size := 8
	if len(args) > 1 {
		n, err := number.Number(args[0])
		if err != nil {
			return nil, err
		}
		size = int(n)
	}
	return make(map[string]interface{}, size), nil
}

func inTableSet(args []interface{}) (interface{}, error) {
	if len(args) < 3 {
		return nil, consts.ErrArgsNotEnough(args)
	}

	var (
		v   interface{}
		val interface{}
	)

	val = args[2]
	v, ok := number.Value(val)
	if ok {
		return table.Set(args[0], args[1], v)
	}

	v, ok = boolean.Value(val)
	if ok {
		return table.Set(args[0], args[1], v)
	}

	v, ok = val.(*string)
	if ok {
		return table.Set(args[0], args[1], v)
	}

	return table.Set(args[0], args[1], val)
}

func inTableGet(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, nil
	}
	v, err := table.Get(args[0], args[1])
	if err != nil {
		return nil, err
	}

	// return default value
	if v == nil && len(args) > 2 {
		v = args[2]
	}
	return v, nil
}

func inTableDel(args []interface{}) (interface{}, error) {
	if len(args) < 2 {
		return nil, nil
	}
	return table.Del(args[0], args[1])
}

func inTableSetInt(args []interface{}) (interface{}, error) {
	if len(args) < 3 {
		return nil, consts.ErrArgsNotEnough(args)
	}
	var (
		val interface{}
		err error
	)

	val = args[2]
	_, ok := number.Value(val)
	if ok {
		var fv float64
		fv, err = number.Number(val)
		val = int64(fv)
	}
	if err != nil {
		return nil, err
	}
	return table.Set(args[0], args[1], val)
}
