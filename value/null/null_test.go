package null

import (
	"go/token"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Type(t *testing.T) {
	tests := []struct {
		val    interface{}
		result bool
	}{
		{nil, true},
		{12, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.result, Type(tt.val))
	}
}

func Test_Eval(t *testing.T) {
	var mmp map[string]interface{}
	var sp *struct {
		Test string
	}

	spp := &struct {
		Str string
	}{}

	tests := []struct {
		left, right interface{}
		op          token.Token
		result      bool
		wantErr     bool
	}{
		{nil, nil, token.EQL, true, false},
		{nil, nil, token.NEQ, false, false},
		{nil, nil, token.ADD, false, true},
		{12, nil, token.EQL, false, false},
		{nil, "test", token.NEQ, true, false},
		{mmp, nil, token.EQL, true, false},
		{nil, sp, token.EQL, true, false},
		{spp, nil, token.NEQ, true, false},
	}

	for _, tt := range tests {
		r, e := Eval(tt.left, tt.right, tt.op)
		assert.Equal(t, tt.result, r.(bool))
		assert.Equal(t, tt.wantErr, e != nil)
	}
}

func Benchmark_EvalPtr(b *testing.B) {
	spp := &struct {
		Str string
	}{}
	b.ResetTimer()
	now := time.Now()
	var (
		rst interface{}
		err error
	)
	for n := 0; n < b.N; n++ {
		rst, err = Eval(spp, nil, token.NEQ)
	}
	assert.Equal(b, true, rst)
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
}

func Benchmark_EvalNil(b *testing.B) {
	var (
		spp interface{}
		rst interface{}
		err error
	)
	b.ResetTimer()
	now := time.Now()
	for n := 0; n < b.N; n++ {
		rst, err = Eval(spp, nil, token.NEQ)
	}
	assert.Equal(b, false, rst)
	assert.Nil(b, err)
	b.Logf("N: %v, cost: %v", b.N, time.Since(now))
}
