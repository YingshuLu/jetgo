package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFor_Completed(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			i = 1
			for {
				if i >= 12 {
					break
				}

				if i == 10 {
					i = i + 1
					continue
				}

				sum = sum + i
				i = i + 1

				for j = 0; j < 10; j = j + 1 {
					if j == 1 {
						sum = sum + 1
						break
					}
				}
				
			}
			return sum
		}
`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	val, err := v.Run(rule, nil)
	assert.Nil(t, err)
	t.Log(val)
	assert.Equal(t, float64(66), val)
}

func TestFor_NoInit(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			i = 1
			for ; i < 12; i = i + 1 {
				if i >= 12 {
					break
				}
				sum = sum + i
			}
			return sum
		}
`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	val, err := v.Run(rule, nil)
	assert.Nil(t, err)
	t.Log(val)
	assert.Equal(t, float64(66), val)
}

func TestFor_NoCond(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			for i = 1;; i = i + 1 {
				if i >= 12 {
					break
				}
				sum = sum + i
			}
			return sum
		}
`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	val, err := v.Run(rule, nil)
	assert.Nil(t, err)
	t.Log(val)
	assert.Equal(t, float64(66), val)
}

func TestFor_NoPost(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			for i = 1;;{
				if i >= 12 {
					break
				}
				sum = sum + i
				i = i + 1
			}
			return sum
		}
`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	val, err := v.Run(rule, nil)
	assert.Nil(t, err)
	t.Log(val)
	assert.Equal(t, float64(66), val)
}

func TestFor_Return(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			for i = 1;;{
				if i == 10 {
					return i
				}
				sum = sum + i
				i = i + 1
			}
			return sum
		}
`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	v := New()

	val, err := v.Run(rule, nil)
	assert.Nil(t, err)
	t.Log(val)
	assert.Equal(t, float64(10), val)
}

func TestFor_Err(t *testing.T) {
	expr := `
		rule  for_cond {
			sum = 0
			for i = 1; i = i + 1; i = i + 1 {
				if i >= 12 {
					break
				}
				sum = sum + i
			}
			return sum
		}
`
	_, err := Compile(expr)
	assert.NotNil(t, err)
}
