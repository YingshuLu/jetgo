// Package array value of expression
package array

import (
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
)

// Type check if arr is array
func Type(any interface{}) bool {
	switch any.(type) {
	case []float32, []float64:
		return true
	case []bool, []string:
		return true
	case []int8, []int16, []int32, []int64, []int:
		return true
	case []uint8, []uint16, []uint32, []uint64, []uint:
		return true
	case []*float32, []*float64:
		return true
	case []*bool, []*string:
		return true
	case []*int8, []*int16, []*int32, []*int64, []*int:
		return true
	case []*uint8, []*uint16, []*uint32, []*uint64, []*uint:
		return true
	case []interface{}:
		return true
	default:
		return false
	}
}

// In check if any in slice
func In(any interface{}, slice interface{}) (interface{}, error) {
	if number.Type(any) {
		return inNumberArray(any, slice)
	}
	if stringv.Type(any) {
		return inStringArray(any, slice)
	}
	if boolean.Type(any) {
		return inBoolArray(any, slice)
	}
	return inInterfaceSlice(any, slice)
}

// Append any behind slice
func Append(slice interface{}, any interface{}) (interface{}, error) {
	if Type(any) {
		return inAppendSlice(slice, any)
	}
	return inAppendElem(slice, any)
}
