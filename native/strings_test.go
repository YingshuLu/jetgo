package native

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConv(t *testing.T) {
	tests := []struct {
		name    string
		str     interface{}
		up      interface{}
		low     interface{}
		wantErr bool
	}{
		{
			name:    "lower string",
			str:     "abc",
			up:      "ABC",
			low:     "abc",
			wantErr: false,
		},
		{
			name:    "upper string",
			str:     "ABC",
			up:      "ABC",
			low:     "abc",
			wantErr: false,
		},
		{
			name:    "lower bytes",
			str:     []byte("abc"),
			up:      "ABC",
			low:     "abc",
			wantErr: false,
		},
		{
			name:    "upper bytes",
			str:     []byte("ABC"),
			up:      "ABC",
			low:     "abc",
			wantErr: false,
		},
		{
			name:    "illegal",
			str:     123,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uv, err := toSupper([]interface{}{tt.str})
			assert.Equal(t, tt.up, uv)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v wantErr: %v error: %v", tt.name, tt.wantErr, err)
			}

			lv, err := toSlower([]interface{}{tt.str})
			assert.Equal(t, tt.low, lv)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v wantErr: %v error: %v", tt.name, tt.wantErr, err)
			}
		})
	}
}

func TestToSplit(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		wantErr bool
	}{
		{
			name: "normal call",
			args: []interface{}{
				"tencent@alibaba@bytes@",
				"@",
			},
			wantErr: false,
		},
		{
			name: "target not string",
			args: []interface{}{
				123,
				"@",
			},
			wantErr: true,
		},
		{
			name: "sp not string",
			args: []interface{}{
				"tencent@alibaba@bytes",
				123,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := inSplit(tt.args)
			t.Logf("args: %v, result: %v", tt.args, v)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v, args: %v, wantErr: %v, err: %v",
					tt.name, tt.args, tt.wantErr, err)
			}
		})
	}
}
