package util

import (
	"fmt"

	"github.com/nicholasq/glox/ast"
	"github.com/nicholasq/glox/token"
)

type AstPrinter struct{}

func (aP *AstPrinter) VisitBinaryExpr(expr *ast.Binary) interface{} {
	return aP.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (aP *AstPrinter) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	return aP.parenthesize("group", expr.Expression)
}

func (aP *AstPrinter) VisitLiteralExpr(expr *ast.Literal) interface{} {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (aP *AstPrinter) VisitUnaryExpr(expr *ast.Unary) interface{} {
	return aP.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (aP *AstPrinter) VisitVariableExpr(expr *ast.Variable) interface{} {
	return expr
}

func (aP *AstPrinter) parenthesize(name string, exprs ...ast.Expr) string {
	result := fmt.Sprintf("(%s", name)
	for _, expr := range exprs {
		result += " "
		result += expr.Accept(aP).(string)
	}
	result += ")"
	return result
}

func (aP *AstPrinter) Print(expr ast.Expr) string {
	return expr.Accept(aP).(string)
}

func TestPrinter() {
	unary := ast.Unary{
		Operator: token.Token{Lexeme: "-", Literal: nil, TokenType: token.MINUS, Line: 1},
		Right:    &ast.Literal{Value: 123},
	}
	operator := token.Token{
		Lexeme: "*", Literal: nil, TokenType: token.STAR, Line: 1,
	}
	grouping := ast.Grouping{
		Expression: &ast.Binary{
			Left:     &ast.Literal{Value: 45.67},
			Operator: token.Token{Lexeme: "+", Literal: nil, TokenType: token.PLUS, Line: 1},
			Right:    &ast.Literal{Value: 78.9},
		},
	}

	expr := &ast.Binary{
		Left:     &unary,
		Operator: operator,
		Right:    &grouping,
	}

	printer := AstPrinter{}
	fmt.Println(printer.Print(expr))
}
