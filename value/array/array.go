package array

import (
	"fmt"
	"go/token"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
)

// Get value of array
func Get(any interface{}, idx int) (interface{}, error) {
	switch slice := any.(type) {
	case []float32:
		return slice[idx], nil
	case []float64:
		return slice[idx], nil
	case []bool:
		return slice[idx], nil
	case []string:
		return slice[idx], nil
	case []int8:
		return slice[idx], nil
	case []int16:
		return slice[idx], nil
	case []int32:
		return slice[idx], nil
	case []int64:
		return slice[idx], nil
	case []int:
		return slice[idx], nil
	case []uint8:
		return slice[idx], nil
	case []uint16:
		return slice[idx], nil
	case []uint32:
		return slice[idx], nil
	case []uint64:
		return slice[idx], nil
	case []uint:
		return slice[idx], nil
	case []*float32:
		return slice[idx], nil
	case []*float64:
		return slice[idx], nil
	case []*bool:
		return slice[idx], nil
	case []*string:
		return slice[idx], nil
	case []*int8:
		return slice[idx], nil
	case []*int16:
		return slice[idx], nil
	case []*int32:
		return slice[idx], nil
	case []*int64:
		return slice[idx], nil
	case []*int:
		return slice[idx], nil
	case []*uint8:
		return slice[idx], nil
	case []*uint16:
		return slice[idx], nil
	case []*uint32:
		return slice[idx], nil
	case []*uint64:
		return slice[idx], nil
	case []*uint:
		return slice[idx], nil
	case []interface{}:
		return slice[idx], nil
	default:
		return nil, consts.ErrNotSlice(any)
	}
}

// Len of array
func Len(any interface{}) (int, error) {
	switch slice := any.(type) {
	case []float32:
		return len(slice), nil
	case []float64:
		return len(slice), nil
	case []bool:
		return len(slice), nil
	case []string:
		return len(slice), nil
	case []int8:
		return len(slice), nil
	case []int16:
		return len(slice), nil
	case []int32:
		return len(slice), nil
	case []int64:
		return len(slice), nil
	case []int:
		return len(slice), nil
	case []uint8:
		return len(slice), nil
	case []uint16:
		return len(slice), nil
	case []uint32:
		return len(slice), nil
	case []uint64:
		return len(slice), nil
	case []uint:
		return len(slice), nil
	case []*float32:
		return len(slice), nil
	case []*float64:
		return len(slice), nil
	case []*bool:
		return len(slice), nil
	case []*string:
		return len(slice), nil
	case []*int8:
		return len(slice), nil
	case []*int16:
		return len(slice), nil
	case []*int32:
		return len(slice), nil
	case []*int64:
		return len(slice), nil
	case []*int:
		return len(slice), nil
	case []*uint8:
		return len(slice), nil
	case []*uint16:
		return len(slice), nil
	case []*uint32:
		return len(slice), nil
	case []*uint64:
		return len(slice), nil
	case []*uint:
		return len(slice), nil
	case []interface{}:
		return len(slice), nil
	default:
		return 0, consts.ErrNotSlice(any)
	}
}

func inNumberArray(any, slice interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)

	switch arr := slice.(type) {
	case []float32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []float64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []int8:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []int16:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []int32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []int64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []int:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}

	case []uint8:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []uint16:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []uint32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []uint64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []uint:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}

	case []*float32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*float64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}

	case []*int8:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*int16:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*int32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*int64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*int:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*uint8:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*uint16:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*uint32:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*uint64:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*uint:
		for _, v := range arr {
			res, err = number.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	default:
		err = consts.ErrUndefinedType(fmt.Sprintf("slice not number array, but: %T", slice))
	}
	return false, err
}

func inStringArray(any, slice interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)

	switch arr := slice.(type) {
	case []string:
		for _, v := range arr {
			res, err = stringv.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*string:
		for _, v := range arr {
			res, err = stringv.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	default:
		err = consts.ErrUndefinedType(fmt.Sprintf("slice not string array, but: %T", slice))
	}

	return false, err
}

func inBoolArray(any, slice interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)

	switch arr := slice.(type) {
	case []bool:
		for _, v := range arr {
			res, err = boolean.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	case []*bool:
		for _, v := range arr {
			res, err = boolean.Eval(any, v, token.EQL)
			if err == nil && res.(bool) {
				return true, nil
			}
			if err != nil {
				return false, err
			}
		}
	default:
		err = consts.ErrUndefinedType(fmt.Sprintf("slice not bool array, but: %T", slice))
	}

	return false, err
}

func inInterfaceSlice(any, slice interface{}) (interface{}, error) {
	arr, ok := slice.([]interface{})
	if !ok {
		return nil, consts.ErrNotSlice(slice)
	}
	for _, v := range arr {
		if v == any {
			return true, nil
		}
	}
	return false, nil
}

func inAppendElem(arr, elem interface{}) (interface{}, error) {
	n, err := Len(arr)
	if err != nil {
		return nil, err
	}

	na := make([]interface{}, 0, n+1)
	var e interface{}
	for i := 0; i < n; i++ {
		e, err = Get(arr, i)
		if err != nil {
			return nil, err
		}
		na = append(na, e)
	}
	return append(na, elem), nil
}

func inAppendSlice(arr, slice interface{}) (interface{}, error) {
	n, err := Len(arr)
	if err != nil {
		return nil, err
	}

	m, err := Len(slice)
	if err != nil {
		return nil, err
	}

	na := make([]interface{}, 0, n+m)
	var e interface{}
	for i := 0; i < n; i++ {
		e, err = Get(arr, i)
		if err != nil {
			return nil, err
		}
		na = append(na, e)
	}

	for j := 0; j < m; j++ {
		e, err = Get(slice, j)
		if err != nil {
			return nil, err
		}
		na = append(na, e)
	}
	return na, nil
}
