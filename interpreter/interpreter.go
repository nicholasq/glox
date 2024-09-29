package interpreter

import (
	"fmt"

	"github.com/nicholasq/glox/ast"
	"github.com/nicholasq/glox/token"
)

type Interpreter struct {
	globals     interface{}
	environment interface{}
	locals      map[interface{}]int
}

func (i *Interpreter) Interpret(statements []ast.Stmt) error {
	// todo figure out a way to return errors from evaluate
	for _, stmt := range statements {
		i.execute(stmt)
	}
	return nil
}

func (i *Interpreter) execute(stmt ast.Stmt) {
	stmt.Accept(i)
}

func stringify(value interface{}) string {
	if value == nil {
		return "nil"
	}
	if flt, ok := value.(float64); ok {
		fltStr := fmt.Sprintf("%v", flt)
		return fltStr
	}

	return fmt.Sprintf("%v", value)
}

func (i *Interpreter) VisitBinaryExpr(expr *ast.Binary) interface{} {
	left := i.evaluate(expr.Left)
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case token.MINUS:
		return left.(float64) - right.(float64)
	case token.SLASH:
		return left.(float64) / right.(float64)
	case token.STAR:
		return left.(float64) * right.(float64)
	case token.PLUS: // todo: add string concatenation
		return left.(float64) + right.(float64)
	case token.GREATER:
		return left.(float64) > right.(float64)
	case token.GREATER_EQUAL:
		return left.(float64) >= right.(float64)
	case token.LESS:
		return left.(float64) < right.(float64)
	case token.LESS_EQUAL:
		return left.(float64) <= right.(float64)
	case token.BANG_EQUAL:
		return !isEqual(left, right)
	case token.EQUAL_EQUAL:
		return isEqual(left, right)
	}
	return nil
}

func (i *Interpreter) VisitGroupingExpr(expr *ast.Grouping) interface{} {
	return i.evaluate(expr.Expression)
}

func (i *Interpreter) VisitLiteralExpr(expr *ast.Literal) interface{} {
	return expr.Value
}

func (i *Interpreter) VisitUnaryExpr(expr *ast.Unary) interface{} {
	right := i.evaluate(expr.Right)

	switch expr.Operator.TokenType {
	case token.MINUS:
		return -(right.(float64))
	case token.BANG:
		return isTruthy(right)
	default:
		return nil
	}
}

func (i *Interpreter) VisitVariableExpr(expr *ast.Variable) interface{}

func (i *Interpreter) VisitExpressionStmt(stmt *ast.ExpressionStmt) {
	i.evaluate(stmt.Expression)
}

func (i *Interpreter) VisitPrintStmt(stmt *ast.PrintStmt) {
	value := i.evaluate(stmt.Expression)
	strValue := stringify(value)
	fmt.Printf("%v\n", strValue)
}

func (i *Interpreter) VisitVarStmt(stmt *ast.VarStmt) {

}

func (i *Interpreter) evaluate(expr ast.Expr) interface{} {
	return expr.Accept(i)
}

func isTruthy(v interface{}) bool {
	if v == nil {
		return false
	}
	truth, ok := v.(bool)
	if ok {
		return truth
	} else {
		return true
	}
}

func checkNumberOperand(v interface{}) float64 {
	num, ok := v.(float64)
	if ok {
		return num
	} else {
		panic("operand must be a number")
	}
}

// todo: figure out how to compare the many different possible data types
func isEqual(a interface{}, b interface{}) bool {

	if a == nil && b == nil {
		return true
	}
	if a == nil {
		return false
	}
	return false
}
