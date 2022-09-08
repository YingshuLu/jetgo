package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexVar(t *testing.T) {
	var b byte
	for {
		name := varRenaming(b)
		idx := indexing(name)
		assert.Equal(t, b, byte(idx))
		if b == 255 {
			break
		}
		b++
	}
}

func TestRenamed(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"x1__", true},
		{"x2__", true},
		{"X2__", false},
		{"x3_", false},
		{"test", false},
	}
	for _, tt := range tests {
		if got := varRenamed(tt.name); got != tt.want {
			t.Errorf("Renamed() = %v, want %v", got, tt.want)
		}
	}
}

func TestShouldLocalVar(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test", false},
		{"test", true},
		{"rR", true},
	}
	for _, tt := range tests {
		if got := varShouldLocal(tt.name); got != tt.want {
			t.Errorf("ShouldLocalVar() = %v, want %v", got, tt.want)
		}
	}
}

func TestHexConv(t *testing.T) {
	tests := []struct {
		name    string
		hexStr  string
		hex     interface{}
		wantErr bool
	}{
		{
			name:    "lower hex string",
			hexStr:  "0xff",
			hex:     int64(255),
			wantErr: false,
		},
		{
			name:    "upper hex string",
			hexStr:  "0XFF",
			hex:     int64(255),
			wantErr: false,
		},
		{
			name:    "mixed hex string 1",
			hexStr:  "0xFF",
			hex:     int64(255),
			wantErr: false,
		},
		{
			name:    "mixed hex string 2",
			hexStr:  "0Xff",
			hex:     int64(255),
			wantErr: false,
		},
		{
			name:    "illegal hex string",
			hexStr:  "0xTT",
			hex:     int64(0),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := stringIntNumber(tt.hexStr)
			assert.Equal(t, tt.hex, v)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v, wantErr: %v, error: %v", tt.name, tt.wantErr, err)
			}
		})
	}
}
