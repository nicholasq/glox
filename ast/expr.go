package ast

import "github.com/nicholasq/glox/token"

// Expr is an interface representing an expression in the Lox language.
// It defines a single method Accept that implements the Visitor pattern.
type Expr interface {
	Accept(v ExpressionVisitor) interface{}
}

// ExpressionVisitor is an interface that defines methods for visiting different types of expressions.
// It is used to implement the Visitor pattern for traversing and operating on the abstract syntax tree.
type ExpressionVisitor interface {
	VisitBinaryExpr(expr *Binary) interface{}
	VisitGroupingExpr(expr *Grouping) interface{}
	VisitLiteralExpr(expr *Literal) interface{}
	VisitUnaryExpr(expr *Unary) interface{}
	VisitVariableExpr(expr *Variable) interface{}
}

// Binary represents a binary expression in the Lox language.
// It consists of a left operand, an operator, and a right operand.
type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func (expr *Binary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitBinaryExpr(expr)
}

// Grouping represents a grouping expression in the Lox language.
// It contains a single expression that is enclosed in parentheses.
type Grouping struct {
	Expression Expr
}

func (expr *Grouping) Accept(v ExpressionVisitor) interface{} {
	return v.VisitGroupingExpr(expr)
}

// Literal represents a literal value in the Lox language.
// It can hold various types of values such as numbers, strings, or booleans.
type Literal struct {
	Value interface{}
}

func (expr *Literal) Accept(v ExpressionVisitor) interface{} {
	return v.VisitLiteralExpr(expr)
}

// Unary represents a unary expression in the Lox language.
// It consists of an operator token and a right operand expression.
type Unary struct {
	Operator token.Token
	Right    Expr
}

func (expr *Unary) Accept(v ExpressionVisitor) interface{} {
	return v.VisitUnaryExpr(expr)
}

// Variable represents a variable expression in the Lox language.
// It contains a token that holds the name of the variable.
type Variable struct {
	Name token.Token
}

func (expr *Variable) Accept(v ExpressionVisitor) interface{} {
	return v.VisitVariableExpr(expr)
}
