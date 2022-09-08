package vm

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/yingshulu/jetgo/consts"
)

const (
	lparenToken = "{"
	ruleToken   = "rule"
	funcToken   = "func()"
	packageFake = "package rengo\n"
	funcFake    = "func rengo()"
)

type File = *token.FileSet

// Compile source code into Rule
func Compile(source string) (Rule, error) {
	r := &rule{
		source: source,
	}

	s1, err := r.format(source, funcToken)
	if err != nil {
		return nil, err
	}
	expr, err := parser.ParseExpr(s1)
	if err != nil {
		return nil, err
	}
	r.expr = expr

	// compile
	err = r.compile()
	if err != nil {
		return nil, err
	}

	s2, err := r.format(source, funcFake)
	if err != nil {
		return nil, err
	}
	r.file = token.NewFileSet()
	_, err = parser.ParseFile(r.file, r.Name(), packageFake+s2, 0)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Rule interface of compiled source code
type Rule interface {
	Name() string
	Attr() []string
	Source() string
	Expr() ast.Expr
	CallArgs() int
	SymbolTable() map[string]string
	Consts() []interface{}
	File() File
}

type rule struct {
	expr        ast.Expr
	name        string
	source      string
	attrs       []string
	maxCallArgs int
	symbolTable map[string]string
	consts      []interface{}
	file        File
}

// Name of the rule
func (r *rule) Name() string {
	return r.name
}

// Attr attribute of rule defined
func (r *rule) Attr() []string {
	return r.attrs
}

// Source code of the rule
func (r *rule) Source() string {
	return r.source
}

// Expr ast of the rule
func (r *rule) Expr() ast.Expr {
	return r.expr
}

// CallArgs count of call function's args
func (r *rule) CallArgs() int {
	return r.maxCallArgs
}

// SymbolTable of the rule
func (r *rule) SymbolTable() map[string]string {
	return r.symbolTable
}

// Consts of the rule
func (r *rule) Consts() []interface{} {
	return r.consts
}

// File info of the rule
func (r *rule) File() File {
	return r.file
}

func (r *rule) format(source string, prefix string) (string, error) {
	end := strings.Index(source, lparenToken)
	if end == -1 {
		return "", consts.ErrSemanticUndefined("not found {")
	}

	statement := strings.TrimSpace(source[0:end])
	block := source[end:]

	argv := strings.Fields(statement)
	argn := len(argv)
	if argn < 1 || argv[0] != ruleToken {
		return "", consts.ErrSemanticUndefined("not found \"rule\" start token")
	}
	if argn < 2 {
		return "", consts.ErrSemanticUndefined("not assign rule name")
	}
	r.name = argv[1]
	if argn >= 3 {
		r.attrs = argv[2:]
	}
	var builder strings.Builder
	_, err := builder.WriteString(prefix)
	if err != nil {
		return "", err
	}
	_, err = builder.WriteString(" ")
	if err != nil {
		return "", err
	}
	_, err = builder.WriteString(block)
	if err != nil {
		return "", err
	}
	return builder.String(), nil
}
