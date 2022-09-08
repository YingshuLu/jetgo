// Package vm of rule running env
package vm

import (
	"fmt"
	"go/token"
)

// VM code running env
type VM interface {
	Run(Rule, interface{}) (interface{}, error)
}

// New create new VM instance
func New() VM {
	return &vm{
		symbolStack:   make(stack, 16),
		callArgsStack: make(stack, 8),
	}
}

// stack rule running stack memory
type stack []interface{}

func (s stack) growLen(expect int) int {
	n := len(s)
	if expect <= n {
		return n
	}
	v := expect
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func (s stack) growNoCopy(expect int) stack {
	expect = s.growLen(expect)
	if expect <= len(s) {
		return s
	}
	return make(stack, expect)
}

type vm struct {
	symbolStack   stack
	callArgsStack stack
	rule          Rule
	err           error
	branch        token.Token
}

// Run rule with fact
func (v *vm) Run(rule Rule, fact interface{}) (res interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("panic: %v, vm: %s", e, v.debugInfo())
		}
	}()
	v.prepare(rule)
	return v.evalExpr(rule.Expr(), fact)
}

func (v *vm) prepare(r Rule) {
	n := len(r.SymbolTable())
	v.symbolStack = v.symbolStack.growNoCopy(n)
	n = r.CallArgs()
	v.callArgsStack = v.callArgsStack.growNoCopy(n)
	v.rule = r
	v.err = nil
	v.branch = token.ILLEGAL
}
