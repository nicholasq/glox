package main

type Expr interface {
	Accept(v ExpressionVisitor) interface{}
}

type ExpressionVisitor interface {
	VisitBinaryExpr(expr Binary) interface{}
	VisitGroupingExpr(expr Grouping) interface{}
	VisitLiteralExpr(expr Literal) interface{}
	VisitUnaryExpr(expr Unary) interface{}
}

type Binary struct {
	left     Expr
	operator Token
	right    Expr
}

type Grouping struct {
	expression Expr
}

type Literal struct {
	value interface{}
}

type Unary struct {
	operator Token
	right    Expr
}

func (expr Binary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitBinaryExpr(expr)
}

func (expr Grouping) Accept(v ExpressionVisitor) interface{} {
	return v.VisitGroupingExpr(expr)
}

func (expr Literal) Accept(v ExpressionVisitor) interface{} {
	return v.VisitLiteralExpr(expr)
}

func (expr Unary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitUnaryExpr(expr)
}
