package vm

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_returnStmt(t *testing.T) {
	node := &ast.ReturnStmt{}
	v := &vm{}
	val, err := v.returnStmt(node, nil)
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func Test_unaryStmt(t *testing.T) {
	source := `
		rule unary {
			a = -1
			b = -a
			if !C {
				return -D
			}
			return a + b
		}
	`
	r, err := Compile(source)
	assert.Nil(t, err)

	data := map[string]interface{}{
		"C": false,
		"D": -1,
	}
	v := New()
	val, err := v.Run(r, data)
	assert.Nil(t, err)
	assert.Equal(t, val, float64(1))

	data["C"] = true
	val, err = v.Run(r, data)
	assert.Nil(t, err)
	assert.Equal(t, val, float64(0))

	data["C"] = 1
	_, err = v.Run(r, data)
	assert.NotNil(t, err)

	data["C"] = false
	data["D"] = "-1"
	_, err = v.Run(r, data)
	assert.NotNil(t, err)
}

func Test_line(t *testing.T) {
	source := `
		rule kamlu_rule {
			a = -1
			b = -a

			c = 1
			if !C && F == 1 {
				a = a + b
			}

			c = 1 / 0
			return a + b
		}
	`
	r, err := Compile(source)
	if err != nil {
		t.Logf("compile error: %v", err)
	}

	data := map[string]interface{}{
		"C": false,
		"D": -1,
	}
	v := New()
	_, err = v.Run(r, data)
	t.Log(err)
}
