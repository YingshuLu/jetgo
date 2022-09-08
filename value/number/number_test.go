package number

import (
	"go/token"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func Test_Type(t *testing.T) {
	tests := []struct {
		name      string
		val       interface{}
		notNumber bool
	}{
		{
			name: "int",
			val:  12,
		},
		{
			name: "uint",
			val:  uint(13),
		},
		{
			name: "float32",
			val:  float32(16),
		},
		{
			name: "*int32",
			val:  proto.Int32(17),
		},
		{
			name: "*uint32",
			val:  proto.Uint32(18),
		},
		{
			name: "*float64",
			val:  proto.Float64(19.0),
		},
		{
			name:      "not number",
			val:       "123",
			notNumber: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Type(tt.val)
			if tt.notNumber == r {
				t.Logf("name: %v, val type: %T, notNumber: %v",
					tt.name, tt.val, tt.notNumber)
			}
		})
	}
}

func Test_Number(t *testing.T) {
	pint := 1
	puint := uint(2)
	pint8 := int8(3)
	puint8 := uint8(4)
	pint16 := int16(5)
	puint16 := uint16(6)
	pint32 := int32(7)
	puint32 := uint32(8)
	pint64 := int64(9)
	puint64 := uint64(10)
	pfloat32 := float32(11)
	pfloat64 := float64(12)

	values := []interface{}{
		pint,
		puint,
		pint8,
		puint8,
		pint16,
		puint16,
		pint32,
		puint32,
		pint64,
		puint64,
		pfloat32,
		pfloat64,
		&pint,
		&puint,
		&pint8,
		&puint8,
		&pint16,
		&puint16,
		&pint32,
		&puint32,
		&pint64,
		&puint64,
		&pfloat32,
		&pfloat64,
	}

	for _, v := range values {
		_, e := Number(v)
		assert.Nil(t, e)
	}

	_, e := Number("test")
	assert.NotNil(t, e)
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
		result      interface{}
		wantErr     bool
	}{
		{nil, nil, token.EQL, nil, true},
		{12, nil, token.EQL, nil, true},
		{mmp, nil, token.EQL, nil, true},
		{nil, sp, token.EQL, nil, true},
		{spp, nil, token.NEQ, nil, true},
		{3, 2, token.ADD, 5.0, false},
		{3, 2, token.SUB, 1.0, false},
		{2, 3, token.SUB, -1.0, false},
		{3, 2, token.MUL, 6.0, false},
		{3, 2, token.QUO, 1.5, false},
		{3, 2, token.REM, 1.0, false},
		{1, 3, token.AND, 1.0, false},
		{1, 3, token.OR, 3.0, false},
		{1, 3, token.XOR, 2.0, false},
		{1, 3, token.SHL, 8.0, false},
		{8, 2, token.SHR, 2.0, false},
		{1, 3, token.EQL, false, false},
		{1, 3, token.LSS, true, false},
		{1, 3, token.GTR, false, false},
		{1, 3, token.NEQ, true, false},
		{1, 3, token.LEQ, true, false},
		{1, 3, token.GEQ, false, false},
		{proto.Uint32(12), proto.Uint64(6), token.QUO, 2.0, false},
	}

	for _, tt := range tests {
		r, e := Eval(tt.left, tt.right, tt.op)
		assert.Equal(t, tt.result, r)
		assert.Equal(t, tt.wantErr, e != nil)
	}
}
