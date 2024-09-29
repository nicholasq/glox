package ast

import "github.com/nicholasq/glox/token"

type Expr interface {
	Accept(v ExpressionVisitor) interface{}
}

type ExpressionVisitor interface {
	VisitBinaryExpr(expr *Binary) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
}

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *Binary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitBinaryExpr(expr)
}

type Grouping struct {
	Expression Expr
}

func (expr *Grouping) Accept(v ExpressionVisitor) interface{} {
	return v.VisitGroupingExpr(expr)
}

type Literal struct {
	Value interface{}
}

func (expr *Literal) Accept(v ExpressionVisitor) interface{} {
	return v.VisitLiteralExpr(expr)
}

type Unary struct {
	Operator token.Token
	Right    Expr
}

func (expr *Unary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitUnaryExpr(expr)
}

type Variable struct {
	Name token.Token
}

func (expr *Variable) Accept(v ExpressionVisitor) interface{} {
	return v.VisitVariableExpr(expr)
}
