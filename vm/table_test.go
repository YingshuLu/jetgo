package vm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	rule := `
	rule table_test {
		key = "test"
		val = 12.6
		t = table {
			"number": 123,
			"hello": "world",
			key: val,
		}
		return t
	}
`
	r, err := Compile(rule)
	assert.Nil(t, err)

	v := New()
	m, err := v.Run(r, nil)
	assert.Nil(t, err)
	t.Log(m)
}
