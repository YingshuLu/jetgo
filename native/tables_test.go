package native

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

var testSetCases = []struct {
	name    string
	args    []interface{}
	result  interface{}
	wantErr bool
}{
	{
		name: "normal int",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			123,
		},
		result:  123,
		wantErr: false,
	},
	{
		name: "normal int pointer",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			proto.Int32(123),
		},
		result:  int32(123),
		wantErr: false,
	},
	{
		name: "normal bool pointer",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			true,
		},
		result:  true,
		wantErr: false,
	},
	{
		name: "normal bool pointer",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			proto.Bool(true),
		},
		result:  true,
		wantErr: false,
	},
	{
		name: "normal string",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			"value",
		},
		result:  "value",
		wantErr: false,
	},
	{
		name: "normal string pointer",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
			proto.String("test"),
		},
		result:  "test",
		wantErr: false,
	},
	{
		name: "args not enough",
		args: []interface{}{
			make(map[string]interface{}),
			"test",
		},
		wantErr: true,
	},
}

func TestTableSet(t *testing.T) {
	for _, tt := range testSetCases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := inTableSetInt(tt.args)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v, args: %v, wantErr: %v, err: %v",
					tt.name, tt.args, tt.wantErr, err)
			}
			if err != nil {
				v, err := inTableGet(tt.args)
				assert.Nil(t, err)
				assert.Equal(t, tt.result, v)
			}
		})
	}
}

func TestTableSetInt(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		result  interface{}
		wantErr bool
	}{
		{
			name: "normal int",
			args: []interface{}{
				make(map[string]interface{}),
				"test",
				123,
			},
			result:  123,
			wantErr: false,
		},
		{
			name: "normal float",
			args: []interface{}{
				make(map[string]interface{}),
				"test",
				float64(123),
			},
			result:  float64(123),
			wantErr: false,
		},
		{
			name: "normal string",
			args: []interface{}{
				make(map[string]interface{}),
				"test",
				"value",
			},
			result:  "value",
			wantErr: false,
		},
		{
			name: "args not enough",
			args: []interface{}{
				make(map[string]interface{}),
				"test",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := inTableSetInt(tt.args)
			if tt.wantErr != (err != nil) {
				t.Errorf("case: %v, args: %v, wantErr: %v, err: %v",
					tt.name, tt.args, tt.wantErr, err)
			}
			if err != nil {
				v, err := inTableGet(tt.args)
				assert.Nil(t, err)
				assert.Equal(t, tt.result, v)
			}
		})
	}
}

func TestTableGet(t *testing.T) {
	tests := []struct {
		name    string
		args    []interface{}
		result  interface{}
		wantErr bool
	}{
		{
			name: "normal get",
			args: []interface{}{
				map[string]interface{}{
					"test": 123,
				},
				"test",
			},
			result:  123,
			wantErr: false,
		},
		{
			name: "default get",
			args: []interface{}{
				map[string]interface{}{
					"test": 123,
				},
				"default_key",
				456,
			},
			result:  456,
			wantErr: false,
		},
		{
			name: "args not enough",
			args: []interface{}{
				make(map[string]interface{}),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := inTableGet(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("case: %v, args: %v, wantErr: %v, err: %v",
					tt.name, tt.args, tt.wantErr, err)
			} else {
				assert.Equal(t, tt.result, v)
			}
		})
	}

}
