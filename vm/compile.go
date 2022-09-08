package vm

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/yingshulu/jetgo/consts"
	"github.com/yingshulu/jetgo/native"
	"github.com/yingshulu/jetgo/value"
	"github.com/yingshulu/jetgo/value/boolean"
	"github.com/yingshulu/jetgo/value/null"
	"github.com/yingshulu/jetgo/value/number"
	stringv "github.com/yingshulu/jetgo/value/string"
)

type visitor struct {
	rule        *rule
	err         error
	vindex      byte
	cindex      byte
	maxCallArgs int
	symbolTable map[string]string
	constsTable map[string]string
}

func (v *visitor) basicVal(node *ast.BasicLit) (interface{}, error) {
	// 1. renamed
	if constRenamed(node.Value) {
		index := indexing(node.Value)
		return v.rule.consts[index], nil
	}

	// 2. same consts, use stack value and rename
	name, ok := v.constsTable[node.Value]
	if ok {
		node.Value = name
		index := indexing(name)
		return v.rule.consts[index], nil
	}

	// 3. new consts
	var (
		val interface{}
		err error
	)
	switch node.Kind {
	case token.STRING:
		val, err = stringTrim(node.Value)
	case token.INT:
		val, err = stringIntNumber(node.Value)
	case token.FLOAT:
		val, err = stringFloatNumber(node.Value)
	default:
		err = consts.ErrUndefinedType(node.Kind)
	}
	if err != nil {
		return nil, err
	}

	name = constRenaming(v.cindex)
	v.constsTable[node.Value] = name
	v.rule.consts = append(v.rule.consts, val)
	node.Value = name
	v.cindex++
	return val, err
}

func (v *visitor) identVal(node *ast.Ident) (interface{}, error) {
	switch node.Name {
	case trueToken:
		return true, nil
	case falseToken:
		return false, nil
	case nullToken:
		return nil, nil
	default:
		return nil, consts.ErrVarNotFound(node.Name)
	}
}

func (v *visitor) logicBinaryExpr(xv, yv interface{}, xe, ye error, op token.Token) ast.Expr {
	if op == token.AND {
		if xe == nil {
			xb, ok := boolean.Boolean(xv)
			if ok && !xb {
				return ast.NewIdent(falseToken)
			}
		}

		if ye == nil {
			yb, ok := boolean.Boolean(yv)
			if ok && !yb {
				return ast.NewIdent(falseToken)
			}
		}
	}

	if op == token.OR {
		if xe == nil {
			xb, ok := boolean.Boolean(xv)
			if ok && xb {
				return ast.NewIdent(trueToken)
			}
		}

		if ye == nil {
			yb, ok := boolean.Boolean(yv)
			if ok && yb {
				return ast.NewIdent(trueToken)
			}
		}
	}
	return nil
}

func (v *visitor) numberBinaryExpr(xv, yv interface{}, op token.Token) ast.Expr {
	if !number.Type(xv) || !number.Type(yv) {
		return nil
	}
	val, err := number.Eval(xv, yv, op)
	if err != nil {
		return nil
	}

	if value.Type(val) == value.NumberType {
		nd := &ast.BasicLit{
			Kind:  token.FLOAT,
			Value: fmt.Sprint(val),
		}
		v.basicVal(nd)
		return nd
	}

	nd := &ast.Ident{
		Name: fmt.Sprint(val),
	}
	return nd
}

func (v *visitor) stringBinaryExpr(xv, yv interface{}, op token.Token) ast.Expr {
	if !stringv.Type(xv) || !stringv.Type(yv) {
		return nil
	}

	val, err := stringv.Eval(xv, yv, op)
	if err != nil {
		return nil
	}

	if value.Type(val) == value.StringType {
		nd := &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("\"%v\"", val),
		}
		v.basicVal(nd)
		return nd
	}
	nd := &ast.Ident{
		Name: fmt.Sprint(val),
	}
	return nd
}

func (v *visitor) boolBinaryExpr(xv, yv interface{}, op token.Token) ast.Expr {
	if boolean.Type(xv) && boolean.Type(yv) {
		val, err := boolean.Eval(xv, yv, op)
		if err != nil {
			return nil
		}
		name := falseToken
		if val.(bool) {
			name = trueToken
		}
		return ast.NewIdent(name)
	}
	return nil
}

func (v *visitor) nullBinaryExpr(xv, yv interface{}, op token.Token) ast.Expr {
	if null.Type(xv) || null.Type(yv) {
		val, err := null.Eval(xv, yv, op)
		if err != nil {
			return nil
		}
		name := falseToken
		if val.(bool) {
			name = trueToken
		}
		return &ast.Ident{Name: name}
	}
	return nil
}

func (v *visitor) binaryExpr(node *ast.BinaryExpr) (ast.Expr, error) {
	var (
		xv, yv interface{}
		err    error
	)

	xbn, ok := node.X.(*ast.BinaryExpr)
	if ok {
		node.X, err = v.binaryExpr(xbn)
		if err != nil {
			return node, err
		}
	}

	ybn, ok := node.Y.(*ast.BinaryExpr)
	if ok {
		node.Y, err = v.binaryExpr(ybn)
		if err != nil {
			return node, err
		}
	}

	xv, xe := v.valueOf(node.X)
	yv, ye := v.valueOf(node.Y)

	nd := v.logicBinaryExpr(xv, yv, xe, ye, node.Op)
	if nd != nil {
		return nd, nil
	}

	if xe != nil || ye != nil {
		return node, nil
	}

	nd = v.numberBinaryExpr(xv, yv, node.Op)
	if nd != nil {
		return nd, nil
	}

	nd = v.stringBinaryExpr(xv, yv, node.Op)
	if nd != nil {
		return nd, nil
	}

	nd = v.boolBinaryExpr(xv, yv, node.Op)
	if nd != nil {
		return nd, nil
	}

	nd = v.nullBinaryExpr(xv, yv, node.Op)
	if nd != nil {
		return nd, nil
	}

	return node, nil
}

func (v *visitor) valueOf(node ast.Expr) (interface{}, error) {
	ba, ok := node.(*ast.BasicLit)
	if ok {
		val, err := v.basicVal(ba)
		if err == nil {
			return val, err
		}
	}

	xi, ok := node.(*ast.Ident)
	if ok {
		val, err := v.identVal(xi)
		if err == nil {
			return val, err
		}
	}
	return nil, consts.ErrVarNotFound(nil)
}

func (v *visitor) assignStmt(node *ast.AssignStmt) error {
	lhs, rhs := node.Lhs, node.Rhs

	// only support '=' assign
	if len(lhs) > 1 || len(rhs) > 1 || node.Tok != token.ASSIGN {
		return consts.ErrSemanticUndefined(node.TokPos)
	}
	ident, ok := lhs[0].(*ast.Ident)
	if !ok {
		return consts.ErrSemanticUndefined(node.TokPos)
	}
	name, existed := v.symbolTable[ident.Name]
	if !existed {
		name = varRenaming(v.vindex)
		v.symbolTable[ident.Name] = name
		if v.vindex == 255 {
			return consts.ErrToManyLocalVars
		}
		v.vindex++
	}
	ident.Name = name

	bn, ok := rhs[0].(*ast.BinaryExpr)
	if ok {
		nd, err := v.binaryExpr(bn)
		if err != nil {
			return err
		}
		node.Rhs[0] = nd
	}
	return nil
}

// only support func call expr stmt
func (v *visitor) exprStmt(node *ast.ExprStmt) error {
	_, ok := node.X.(*ast.CallExpr)
	if !ok {
		return consts.ErrSemanticUndefined(node)
	}
	return nil
}

// load consts variants
func (v *visitor) basicLit(node *ast.BasicLit) error {
	_, err := v.basicVal(node)
	return err
}

func (v *visitor) ident(node *ast.Ident) error {
	if !varShouldLocal(node.Name) || varRenamed(node.Name) || varToken(node.Name) {
		return nil
	}

	_, ok := native.Fetch(node.Name)
	if ok {
		return nil
	}

	name, ok := v.symbolTable[node.Name]
	if !ok {
		return consts.ErrVarNotFound(fmt.Sprintf("%v - %v", node.Name, node.Pos()))
	}
	node.Name = name
	return nil
}

func (v *visitor) callExpr(node *ast.CallExpr) error {
	identNode, ok := node.Fun.(*ast.Ident)
	if !ok {
		return consts.ErrSemanticUndefined(node.Fun.Pos())
	}
	name := identNode.Name
	_, ok = native.Fetch(name)
	if !ok {
		return consts.ErrFuncNotFound(fmt.Sprintf("%v - %v", identNode.Name, identNode.Pos()))
	}
	v.getCallExprArgsNum(node)
	return nil
}

func (v *visitor) getCallExprArgsNum(expr ast.Expr) int {
	node, ok := expr.(*ast.CallExpr)
	if !ok {
		return 0
	}

	n := len(node.Args)
	for _, ex := range node.Args {
		n += v.getCallExprArgsNum(ex)
	}

	if n > v.maxCallArgs {
		v.maxCallArgs = n
	}

	return n
}

func (v *visitor) ifStmt(node *ast.IfStmt) error {
	if node.Else != nil {
		return consts.ErrSemanticUndefined(
			fmt.Sprintf("unsupport 'else' pos: %d", node.Else.Pos()))
	}
	bn, ok := node.Cond.(*ast.BinaryExpr)
	if ok {
		nd, err := v.binaryExpr(bn)
		if err != nil {
			return err
		}
		node.Cond = nd
	}
	return nil
}

func (v *visitor) compositeLit(node *ast.CompositeLit) error {
	switch node.Type.(type) {
	case *ast.ArrayType:
		return v.forkCompose(node)
	case *ast.Ident:
		return v.tableCompose(node)
	default:
		return consts.ErrSemanticUndefined(node.Pos())
	}
}

func (v *visitor) forkCompose(node *ast.CompositeLit) error {
	arr, ok := node.Type.(*ast.ArrayType)
	if !ok {
		return consts.ErrSemanticUndefined(node.Type.Pos())
	}
	tk, ok := arr.Elt.(*ast.Ident)
	if !ok {
		return consts.ErrSemanticUndefined(arr.Elt.Pos())
	}
	if tk.Name != forkToken {
		return consts.ErrSemanticUndefined(tk.Name)
	}
	for _, expr := range node.Elts {
		fn, ok := expr.(*ast.CallExpr)
		if !ok {
			return consts.ErrSemanticUndefined(expr.Pos())
		}
		err := v.callExpr(fn)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v *visitor) tableCompose(node *ast.CompositeLit) error {
	indent, ok := node.Type.(*ast.Ident)
	if !ok {
		return consts.ErrSemanticUndefined(node.Type.Pos())
	}
	if indent.Name != tableToken {
		return consts.ErrSemanticUndefined(indent.Name)
	}
	var err error
	for _, expr := range node.Elts {
		kv, ok := expr.(*ast.KeyValueExpr)
		if !ok {
			return consts.ErrSemanticUndefined(expr.Pos())
		}
		v.Visit(kv.Key)
		err = v.err
		if err != nil {
			break
		}
		v.Visit(kv.Value)
		err = v.err
		if err != nil {
			break
		}
	}
	return err
}

// Visit optimize source rule
func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	var err error
	switch nd := node.(type) {
	case *ast.AssignStmt:
		err = v.assignStmt(nd)
	case *ast.Ident:
		err = v.ident(nd)
	case *ast.BasicLit:
		err = v.basicLit(nd)
	case *ast.IfStmt:
		err = v.ifStmt(nd)
	case *ast.CallExpr:
		err = v.callExpr(nd)
	case *ast.ExprStmt:
		err = v.exprStmt(nd)
	case *ast.CompositeLit:
		err = v.compositeLit(nd)
	default:
	}
	if err != nil {
		v.err = err
		return nil
	}
	return v
}

func (r *rule) compile() error {
	v := &visitor{
		rule:        r,
		symbolTable: make(map[string]string),
		constsTable: make(map[string]string),
	}

	ast.Walk(v, r.expr)
	if v.err != nil {
		return v.err
	}
	if len(v.symbolTable) > 0 {
		r.symbolTable = v.symbolTable
	}
	r.maxCallArgs = v.maxCallArgs
	return nil
}
