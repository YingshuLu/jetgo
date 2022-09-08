// Package table value of expression
package table

import (
	"github.com/yingshulu/jetgo/consts"
	stringv "github.com/yingshulu/jetgo/value/string"
)

// Type check if any is table
func Type(any interface{}) bool {
	_, ok := any.(map[string]interface{})
	return ok
}

// Value check if any table, return table if yes
func Value(any interface{}) (map[string]interface{}, bool) {
	m, ok := any.(map[string]interface{})
	return m, ok
}

// Get value with key of table
func Get(any, key interface{}) (interface{}, error) {
	t, ok := any.(map[string]interface{})
	if !ok {
		return nil, consts.ErrNotTable(any)
	}
	k, err := stringv.String(key)
	if err != nil {
		return nil, err
	}
	v, ok := t[k]
	if !ok {
		return nil, nil
	}
	return v, nil
}

// Set value with key of table
func Set(any, key, value interface{}) (interface{}, error) {
	t, ok := any.(map[string]interface{})
	if !ok {
		return nil, consts.ErrNotTable(any)
	}

	k, err := stringv.String(key)
	if err != nil {
		return nil, err
	}
	t[k] = value
	return nil, nil
}

// Del value with key of table
func Del(any, key interface{}) (interface{}, error) {
	t, ok := any.(map[string]interface{})
	if !ok {
		return nil, consts.ErrNotTable(any)
	}

	k, err := stringv.String(key)
	if err != nil {
		return nil, err
	}
	delete(t, k)
	return nil, nil
}

// In check if key is in table
func In(key, any interface{}) (interface{}, error) {
	t, ok := any.(map[string]interface{})
	if !ok {
		return nil, consts.ErrNotTable(any)
	}
	k, err := stringv.String(key)
	if err != nil {
		return false, err
	}
	_, ok = t[k]
	return ok, nil
}
