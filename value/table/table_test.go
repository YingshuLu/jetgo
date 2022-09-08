package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Type(t *testing.T) {
	m := make(map[string]interface{})
	v := Type(m)
	assert.True(t, v)
	_, ok := Value(m)
	assert.Equal(t, v, ok)

	v = Type(v)
	assert.False(t, v)
	_, ok = Value(v)
	assert.Equal(t, v, ok)
}

func Test_Get(t *testing.T) {
	tests := []struct {
		name    string
		any     interface{}
		key     interface{}
		wantErr bool
	}{
		{
			name: "has key",
			any: map[string]interface{}{
				"test": 123,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name: "no key",
			any: map[string]interface{}{
				"pojo": 456,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name:    "not map error",
			any:     []string{"good", "bad"},
			key:     "good",
			wantErr: true,
		},
		{
			name: "key not string",
			any: map[string]interface{}{
				"po": 456,
			},
			key:     123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Get(tt.any, tt.key)
			if tt.wantErr != (err != nil) {
				t.Logf("name: %v, wantErr: %v, error: %v",
					tt.name, tt.wantErr, err)
			}
		})
	}
}

func Test_Set(t *testing.T) {
	tests := []struct {
		name    string
		any     interface{}
		key     interface{}
		value   interface{}
		wantErr bool
	}{
		{
			name: "set key success",
			any: map[string]interface{}{
				"test": 123,
			},
			key:     "test",
			value:   456,
			wantErr: false,
		},
		{
			name: "no key",
			any: map[string]interface{}{
				"pojo": 456,
			},
			key:     "test",
			value:   123,
			wantErr: false,
		},
		{
			name:    "not map error",
			any:     []string{"good", "bad"},
			key:     "good",
			wantErr: true,
		},
		{
			name: "key not string",
			any: map[string]interface{}{
				"po": 456,
			},
			key:     123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Set(tt.any, tt.key, tt.value)
			if tt.wantErr != (err != nil) {
				t.Logf("name: %v, wantErr: %v, error: %v",
					tt.name, tt.wantErr, err)
			}
		})
	}
}

func Test_Del(t *testing.T) {
	tests := []struct {
		name    string
		any     interface{}
		key     interface{}
		wantErr bool
	}{
		{
			name: "has key",
			any: map[string]interface{}{
				"test": 123,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name: "no key",
			any: map[string]interface{}{
				"pojo": 456,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name:    "not map error",
			any:     []string{"good", "bad"},
			key:     "good",
			wantErr: true,
		},
		{
			name: "key not string",
			any: map[string]interface{}{
				"po": 456,
			},
			key:     123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Del(tt.any, tt.key)
			if tt.wantErr != (err != nil) {
				t.Logf("name: %v, wantErr: %v, error: %v",
					tt.name, tt.wantErr, err)
			}
		})
	}
}

func Test_In(t *testing.T) {
	tests := []struct {
		name    string
		any     interface{}
		key     interface{}
		wantErr bool
	}{
		{
			name: "has key",
			any: map[string]interface{}{
				"test": 123,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name: "no key",
			any: map[string]interface{}{
				"pojo": 456,
			},
			key:     "test",
			wantErr: false,
		},
		{
			name:    "not map error",
			any:     []string{"good", "bad"},
			key:     "good",
			wantErr: true,
		},
		{
			name: "key not string",
			any: map[string]interface{}{
				"po": 456,
			},
			key:     123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := In(tt.key, tt.any)
			if tt.wantErr != (err != nil) {
				t.Logf("name: %v, wantErr: %v, error: %v",
					tt.name, tt.wantErr, err)
			}
		})
	}
}
