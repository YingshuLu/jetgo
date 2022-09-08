package number

import "github.com/yingshulu/jetgo/consts"

// Number convert number any to float64
func Number(any interface{}) (float64, error) {
	switch v := any.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case *float32:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *float64:
		if v == nil {
			return emptyNumber, nil
		}
		return *v, nil
	case *int8:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *int16:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *int32:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *int64:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *int:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *uint8:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *uint16:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *uint32:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *uint64:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	case *uint:
		if v == nil {
			return emptyNumber, nil
		}
		return float64(*v), nil
	default:
		return emptyNumber, consts.ErrNotNumber(any)
	}
}

// Value check if any is number, and return raw data
func Value(any interface{}) (interface{}, bool) {
	switch v := any.(type) {
	case float32:
		return v, true
	case float64:
		return v, true
	case int8:
		return v, true
	case int16:
		return v, true
	case int32:
		return v, true
	case int64:
		return v, true
	case int:
		return v, true
	case uint8:
		return v, true
	case uint16:
		return v, true
	case uint32:
		return v, true
	case uint64:
		return v, true
	case uint:
		return v, true
	case *float32:
		if v == nil {
			return float32(0.0), true
		}
		return *v, true
	case *float64:
		if v == nil {
			return emptyNumber, true
		}
		return *v, true
	case *int8:
		if v == nil {
			return int8(0), true
		}
		return *v, true
	case *int16:
		if v == nil {
			return int16(0), true
		}
		return *v, true
	case *int32:
		if v == nil {
			return int32(0), true
		}
		return *v, true
	case *int64:
		if v == nil {
			return int64(0), true
		}
		return *v, true
	case *int:
		if v == nil {
			return 0, true
		}
		return *v, true
	case *uint8:
		if v == nil {
			return uint8(0), true
		}
		return *v, true
	case *uint16:
		if v == nil {
			return uint16(0), true
		}
		return *v, true
	case *uint32:
		if v == nil {
			return uint32(0), true
		}
		return *v, true
	case *uint64:
		if v == nil {
			return uint64(0), true
		}
		return *v, true
	case *uint:
		if v == nil {
			return uint(0), true
		}
		return *v, true
	default:
		return emptyNumber, false
	}
}
