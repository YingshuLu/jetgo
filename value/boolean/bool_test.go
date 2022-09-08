package boolean

import (
	"go/token"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_and(t *testing.T) {
	tests := []struct {
		left, right bool
		expect      bool
	}{
		{true, true, true},
		{true, false, false},
		{false, false, false},
	}

	for _, c := range tests {
		assert.Equal(t, c.expect, and(c.left, c.right))
	}
}

func Test_or(t *testing.T) {
	tests := []struct {
		left, right bool
		expect      bool
	}{
		{true, true, true},
		{true, false, true},
		{false, false, false},
	}

	for _, c := range tests {
		assert.Equal(t, c.expect, or(c.left, c.right))
	}
}

func Test_equal(t *testing.T) {
	tests := []struct {
		left, right bool
		expect      bool
	}{
		{true, true, true},
		{true, false, false},
		{false, false, true},
	}

	for _, c := range tests {
		assert.Equal(t, c.expect, equal(c.left, c.right))
	}
}

func Test_notEqual(t *testing.T) {
	tests := []struct {
		left, right bool
		expect      bool
	}{
		{true, true, false},
		{true, false, true},
		{false, false, false},
	}

	for _, c := range tests {
		assert.Equal(t, c.expect, notEqual(c.left, c.right))
	}
}

func TestType(t *testing.T) {
	v := true
	tests := []struct {
		val    interface{}
		expect bool
	}{
		{true, true},
		{false, true},
		{12, false},
		{2.01, false},
		{"test", false},
		{&v, true},
	}
	for _, tt := range tests {
		assert.Equal(t, Type(tt.val), tt.expect)
	}
}

func TestBoolean(t *testing.T) {

	v1 := true
	v2 := false
	tests := []struct {
		val    interface{}
		expect bool
		ok     bool
	}{
		{true, true, true},
		{false, false, true},
		{12, false, false},
		{2.01, false, false},
		{"test", false, false},
		{&v1, true, true},
		{&v2, false, true},
	}
	for _, tt := range tests {
		r, ok := Boolean(tt.val)
		assert.Equal(t, r, tt.expect)
		assert.Equal(t, ok, tt.ok)
	}
}

func TestEvaluate(t *testing.T) {
	type args struct {
		left  interface{}
		right interface{}
		op    token.Token
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "true and true positive",
			args: args{
				left:  true,
				right: true,
				op:    token.LAND,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true and false positive",
			args: args{
				left:  true,
				right: false,
				op:    token.LAND,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "true and int negative",
			args: args{
				left:  true,
				right: 12,
				op:    token.LAND,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "true or true positive",
			args: args{
				left:  true,
				right: true,
				op:    token.LOR,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true or false positive",
			args: args{
				left:  true,
				right: false,
				op:    token.LOR,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true or int negative",
			args: args{
				left:  true,
				right: 12,
				op:    token.LOR,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "true equal true positive",
			args: args{
				left:  true,
				right: true,
				op:    token.EQL,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true equal false positive",
			args: args{
				left:  true,
				right: false,
				op:    token.EQL,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "true equal int negative",
			args: args{
				left:  true,
				right: 12,
				op:    token.EQL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "true not equal true positive",
			args: args{
				left:  true,
				right: true,
				op:    token.NEQ,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "true not equal false positive",
			args: args{
				left:  true,
				right: false,
				op:    token.NEQ,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "true not equal int negative",
			args: args{
				left:  true,
				right: 12,
				op:    token.EQL,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "op illegal",
			args: args{
				left:  true,
				right: false,
				op:    token.ADD,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Eval(tt.args.left, tt.args.right, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Evaluate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
