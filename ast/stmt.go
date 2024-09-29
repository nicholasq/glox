package ast

import "github.com/nicholasq/glox/token"

type Stmt interface {
	Accept(v ExpressionStmtVisitor)
}

type ExpressionStmtVisitor interface {
	VisitExpressionStmt(stmt *ExpressionStmt)
	VisitPrintStmt(stmt *PrintStmt)
	VisitVarStmt(stmt *VarStmt)
}

type ExpressionStmt struct {
	Expression Expr
}

func (expr *ExpressionStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitExpressionStmt(expr)
}

type PrintStmt struct {
	Expression Expr
}

func (expr *PrintStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitPrintStmt(expr)
}

type VarStmt struct {
	Name        token.Token
	Initializer Expr
}

func (expr *VarStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitVarStmt(expr)
}
