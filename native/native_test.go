package native

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMethod(_ []interface{}) (interface{}, error) {
	return nil, nil
}

func Test_Register_Fetch(t *testing.T) {
	Register("test", testMethod)
	_, ok := Fetch("test")
	assert.True(t, ok)
}

func Test_lenMethod(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		result  interface{}
		wantErr bool
	}{
		{
			name:    "string",
			args:    []interface{}{"test"},
			result:  4,
			wantErr: false,
		},
		{
			name:    "array",
			args:    []interface{}{[]int{1, 2, 3}},
			result:  3,
			wantErr: false,
		},
		{
			name: "table",
			args: []interface{}{
				map[string]interface{}{
					"test": 123,
				},
			},
			result:  1,
			wantErr: false,
		},
		{
			name:    "error",
			args:    []interface{}{1, 2, 3},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, e := inLen(tt.args)
			if tt.wantErr != (e != nil) {
				t.Errorf("name: %v, wantErr: %v, error: %v", tt.name, tt.wantErr, e)
			}
			if e == nil {
				assert.Equal(t, tt.result, r)
			}
		})
	}
}

func Test_containString(t *testing.T) {
	tests := []struct {
		name        string
		left, right interface{}
		contain     interface{}
		wantErr     bool
	}{
		{
			name:    "contains",
			left:    "hello world",
			right:   "world",
			contain: true,
		},
		{
			name:    "not contains",
			left:    "hello",
			right:   "world",
			contain: false,
		},
		{
			name:    "left not string",
			left:    123,
			right:   "world",
			contain: false,
			wantErr: true,
		},
		{
			name:    "right not string",
			left:    "hello",
			right:   123,
			contain: false,
			wantErr: true,
		},
		{
			name:    "illegal call",
			contain: false,
		},
	}

	for _, tt := range tests {
		var args []interface{}
		if tt.left != nil {
			args = append(args, tt.left)
		}
		if tt.right != nil {
			args = append(args, tt.right)
		}

		t.Run(tt.name, func(t *testing.T) {
			r, e := containString(args)
			if tt.wantErr != (e != nil) {
				t.Errorf("name: %s, left: %v, right: %v, contain: %v, wantErr: %v, result: %v, error: %v",
					tt.name, tt.left, tt.right, tt.contain, tt.wantErr, r, e)
			}
		})
	}
}

func TestToInteger(t *testing.T) {
	tests := []struct {
		name    string
		val     interface{}
		expect  interface{}
		wantErr bool
	}{
		{
			name:    "float32",
			val:     float32(123),
			expect:  int64(123),
			wantErr: false,
		},
		{
			name:    "float64",
			val:     float64(123),
			expect:  int64(123),
			wantErr: false,
		},
		{
			name:    "int8",
			val:     int8(123),
			expect:  int8(123),
			wantErr: false,
		},
		{
			name:    "string",
			val:     "123",
			expect:  int64(123),
			wantErr: false,
		},
		{
			name:    "bool",
			val:     false,
			expect:  nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := toInteger([]interface{}{tt.val})
			assert.Equal(t, tt.expect, v)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v wantErr: %v error: %v", tt.name, tt.wantErr, err)
			}
		})
	}
}
