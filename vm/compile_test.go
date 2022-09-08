package vm

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/fatih/structs"
	"github.com/stretchr/testify/assert"
)

func Test_compile(t *testing.T) {
	expr := `rule test_rule {
			a = 1
			b = 2
			c = a + b * 2
			if c == 12 {
				return true
			}
			d = 12
			f = 19
			d = 13
			a = array(123, "str", false)
			return false
		}`
	rule, err := Compile(expr)
	assert.Nil(t, err)
	assert.NotNil(t, rule)

	expr1 := `rule {
			c = a + b * 2
			if c == 12 {
				return true
			}
			d = 12
			f = 19
			d = 13
			return false
		}`
	rule, err = Compile(expr1)
	assert.NotNil(t, err)
	assert.Nil(t, rule)

	expr2 := `rule test 123{
			a = 1
			b = 2
			c = a + b * 2
			if c == 12 {
				return true
			}
			d = 12
			f = 19
			d = "13"
			return false
		}`
	rule, err = Compile(expr2)
	assert.Nil(t, err)
	assert.NotNil(t, rule)

	expr3 := `rule test 123{
			a = 1
			b = 4
			c = a + b * 2
			if c == 12 {
				return true
			} 
			d = 12
			f = 19
			d = 13

			in(12, array(d, f))
			return false
		}`
	rule, err = Compile(expr3)
	assert.Nil(t, err)
	assert.NotNil(t, rule)

	rule, err = Compile("   ")
	assert.NotNil(t, err)
	assert.Nil(t, rule)
}

func Test_accelerate(t *testing.T) {
	expr := `
	rule test {
		if true == true && 1 == 1 && 2 == 2 && 3==3 && "test" == "test" && false == false && nil != nil {
			return true
		}
		return false
	}
	`
	rule, err := Compile(expr)
	assert.Nil(t, err)

	fset := token.NewFileSet() // positions are relative to fset
	ast.Print(fset, rule.Expr())
}

func Test_accerlate2(t *testing.T) {
	expr := `
	rule test {
		c = true
		a = c == true && 1 == 3
		b = 1 * 1
		test = "te" + "st"
		good = b + 12
		if true == true && b == 1 && test == "test" && 3 == 2 && good == 13 {
			return true
		}
		return false
	}
	`
	rule, err := Compile(expr)
	assert.Nil(t, err)

	fset := token.NewFileSet() // positions are relative to fset
	ast.Print(fset, rule.Expr())
	v := &vm{}
	t.Log(v.Run(rule, nil))
}

func Test_Structs(t *testing.T) {
	type P struct {
		S *string
		V *int
	}

	type D struct {
		Pd *P
		V  *int
		S  *string
	}

	d := &D{}
	m := structs.Map(d)
	t.Log(m)

	expr := `
		rule test {
			if Fact.Pd == nil {
				return "fact.pd is nil"
			}
			return "fact.pd is not nil"
		}
	`
	r, err := Compile(expr)
	if err != nil {
		return
	}

	t.Log(New().Run(r, Fact(d)))
}
