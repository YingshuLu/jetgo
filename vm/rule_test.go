package vm

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rule(t *testing.T) {
	expr1 := `rule test 123{
			a = 1
			b = 2
			c = a + b * 2
			if c == 12 {
				return true
			}
			d = 12
			f = 19
			e = "test"
			array(123)
		}`

	expr2 := `rule test 123{
			a = 1
			b = 2
			c = a + b * 2
			if c == 12 {
				return true
			} else {
				c = 12
			}
			d = 12
			f = 19
			e = "test"
			array(123)
			return false
		}`

	r, err := Compile(expr2)
	assert.NotNil(t, err)
	assert.Nil(t, r)

	r, err = Compile(expr1)
	assert.Nil(t, err)

	assert.Equal(t, 6, len(r.SymbolTable()))
	assert.Equal(t, "test", r.Name())
	assert.Equal(t, "123", r.Attr()[0])
	assert.True(t, len(r.Source()) > 0)
	assert.NotNil(t, r.Expr())
}

func Test_parse(t *testing.T) {
	expr := `
	func() test {
			a = 1
			b = 2
			c = a + b * 2
			if c == 12 {
				return true
			} else {
				c = 12
			}
			d = 12
			f = 19
			e = "test"
			array(123)
			return false
		}`

	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseExpr(expr)
	if err != nil {
		panic(err)
	}
	// Print the AST.
	ast.Print(fset, f)
}
