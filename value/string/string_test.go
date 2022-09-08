package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var str = "hello"

func TestType(t *testing.T) {
	tests := []struct {
		val    interface{}
		result bool
	}{
		{"test", true},
		{[]byte("test"), true},
		{12, false},
		{&str, true},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Type(tt.val))
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		val     interface{}
		result  string
		wantErr bool
	}{
		{"test", "test", false},
		{[]byte("test"), "test", false},
		{12, "", true},
		{&str, str, false},
	}

	for _, tt := range tests {
		r, err := String(tt.val)
		assert.Equal(t, tt.result, r)
		assert.Equal(t, tt.wantErr, err != nil)
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		val     interface{}
		result  int
		wantErr bool
	}{
		{"test", 4, false},
		{[]byte("test"), 4, false},
		{12, 0, true},
		{&str, len(str), false},
	}

	for _, tt := range tests {
		r, err := Len(tt.val)
		assert.Equal(t, tt.result, r)
		assert.Equal(t, tt.wantErr, err != nil)
	}
}

func TestAdd(t *testing.T) {
	left := str
	right := "world"
	tests := []struct {
		left   string
		right  string
		result string
	}{
		{left, right, left + right},
	}

	for _, tt := range tests {
		r, err := add(tt.left, tt.right)
		assert.Equal(t, tt.result, r)
		assert.Nil(t, err)
	}
}

func TestEqualAndNot(t *testing.T) {
	tests := []struct {
		left   string
		right  string
		result bool
	}{
		{"hello", "hello", true},
		{"world", "word", false},
	}

	for _, tt := range tests {
		r, err := equal(tt.left, tt.right)
		assert.Equal(t, tt.result, r)
		assert.Nil(t, err)

		r, err = notEqual(tt.left, tt.right)
		assert.Equal(t, tt.result, !r.(bool))
		assert.Nil(t, err)
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		target, pattern string
		expect, wantErr bool
	}{
		{"hello_world", "[a-z]*_[a-z]*", true, false},
		{"hello_world", "[a-z]*_[a-z]*", true, false},
		{"hello_world", "hello_gopher", false, false},
		{"hello_world", "*test", false, true},
	}
	for _, tt := range tests {
		v, e := match(tt.target, tt.pattern)
		if tt.wantErr != (e != nil) {
			t.Errorf("target: %s, pattern: %s, expect: %v, wantErr: %v, err: %v",
				tt.target, tt.pattern, tt.expect, tt.wantErr, e)
		}
		assert.Equal(t, tt.expect, v.(bool))
		assert.Equal(t, tt.wantErr, e != nil)
	}
}
