package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestType(t *testing.T) {
	tests := []struct {
		arg  interface{}
		want bool
	}{
		{
			arg:  []int{},
			want: true,
		},
		{
			arg:  []uint{},
			want: true,
		},
		{
			arg:  []*int{},
			want: true,
		},
		{
			arg:  []uint8{},
			want: true,
		},
		{
			arg:  []*uint8{},
			want: true,
		},
		{
			arg:  []bool{},
			want: true,
		},
		{
			arg:  []*bool{},
			want: true,
		},
		{
			arg:  []string{},
			want: true,
		},
		{
			arg:  []*string{},
			want: true,
		},
		{
			arg:  []float64{},
			want: true,
		},
		{
			arg:  []*float64{},
			want: true,
		},
		{
			arg:  []interface{}{},
			want: true,
		},
		{
			arg:  "test",
			want: false,
		},
	}
	for _, tt := range tests {
		if got := Type(tt.arg); got != tt.want {
			t.Errorf("Type() = %v, want %v", got, tt.want)
		}
	}
}

var testSlice = []interface{}{
	[]int{1},
	[]uint{2},
	[]int8{3},
	[]uint8{4},
	[]int16{5},
	[]uint16{6},
	[]int32{7},
	[]uint32{8},
	[]int64{9},
	[]uint64{10},
	[]float32{11},
	[]float64{12},
	[]bool{false},
	[]string{"test"},
	[]*int32{proto.Int32(7)},
	[]*int64{proto.Int64(8)},
	[]*uint32{proto.Uint32(9)},
	[]*uint64{proto.Uint64(10)},
	[]*float32{proto.Float32(11)},
	[]*float64{proto.Float64(12)},
	[]*bool{proto.Bool(false)},
	[]*string{proto.String("test")},
}

func Test_Get(t *testing.T) {
	for _, s := range testSlice {
		_, e := Get(s, 0)
		assert.Nil(t, e)
	}

	p := []struct {
		number int
	}{
		{
			12,
		},
	}
	_, e := Get(p, 0)
	assert.NotNil(t, e)
}

func Test_Len(t *testing.T) {
	for _, s := range testSlice {
		n, e := Len(s)
		assert.Nil(t, e)
		assert.Equal(t, 1, n)
	}
}

func Test_In(t *testing.T) {
	for _, s := range testSlice {
		v, e := Get(s, 0)
		assert.Nil(t, e)
		_, e = In(v, s)
		assert.Nil(t, e)
	}
}

func printArray(arr interface{}, t *testing.T) {
	n, err := Len(arr)
	assert.Nil(t, err)
	for i := 0; i < n; i++ {
		v, err := Get(arr, i)
		assert.Nil(t, err)
		t.Logf("array %v elem: %v", i, v)
	}
}

func Test_Append_OK(t *testing.T) {
	e := 1
	for _, s := range testSlice {
		ss, err := Append(s, e)
		assert.Nil(t, err)
		n, err := Len(ss)
		assert.Nil(t, err)
		v, err := Get(ss, n-1)
		assert.Nil(t, err)
		assert.Equal(t, e, v)
	}

	arr := []interface{}{1, 2, 3}
	for _, s := range testSlice {
		n0, err := Len(s)
		assert.Nil(t, err)

		ss, err := Append(s, arr)
		assert.Nil(t, err)

		n, err := Len(ss)
		assert.Nil(t, err)
		assert.Equal(t, n0+3, n)

		v, err := Get(ss, n-1)
		assert.Nil(t, err)
		assert.Equal(t, 3, v)
		printArray(ss, t)
	}
}

func Test_Append_failure(t *testing.T) {
	any := 12
	e := 12
	_, err := Append(any, e)
	assert.NotNil(t, err)
}
