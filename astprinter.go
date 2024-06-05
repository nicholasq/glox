package main

import "fmt"

type AstPrinter struct{}

func (v AstPrinter) VisitBinaryExpr(expr Binary) interface{} {
	return v.parenthesize(expr.operator.Lexeme, expr.left, expr.right)
}

func (v AstPrinter) VisitGroupingExpr(expr Grouping) interface{} {
	return v.parenthesize("group", expr.expression)
}

func (v AstPrinter) VisitLiteralExpr(expr Literal) interface{} {
	if expr.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.value)
}

func (v AstPrinter) VisitUnaryExpr(expr Unary) interface{} {
	return v.parenthesize(expr.operator.Lexeme, expr.right)
}

func (v AstPrinter) parenthesize(name string, exprs ...Expr) string {
	result := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		result += " "
		result += expr.Accept(v).(string)
	}
	result += ")"
	return result
}

func (v AstPrinter) visitBinaryExpr(expr Binary) interface{} {
	return v.VisitBinaryExpr(expr)
}

func (v AstPrinter) Print(expr Expr) string {
	return expr.Accept(v).(string)
}

func TestPrinter() {
	expr := Binary{
		left:     Unary{operator: Token{Lexeme: "-", Literal: nil, TokenType: MINUS, Line: 1}, right: Literal{value: 123}},
		operator: Token{Lexeme: "*", Literal: nil, TokenType: STAR, Line: 1},
		right:    Grouping{expression: Binary{left: Literal{value: 45.67}, operator: Token{Lexeme: "+", Literal: nil, TokenType: PLUS, Line: 1}, right: Literal{value: 78.9}}},
	}
	printer := AstPrinter{}
	fmt.Println(printer.Print(expr))
}
