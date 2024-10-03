package ast

import "github.com/nicholasq/glox/token"

// Stmt is the interface for all statement types in the AST.
// It defines the Accept method for the visitor pattern.
type Stmt interface {
	Accept(v ExpressionStmtVisitor)
}

// ExpressionStmtVisitor defines the interface for visiting different types of statements.
// Each method corresponds to a specific statement type.
type ExpressionStmtVisitor interface {
	VisitExpressionStmt(stmt *ExpressionStmt)
	VisitPrintStmt(stmt *PrintStmt)
	VisitVarStmt(stmt *VarStmt)
}

// ExpressionStmt represents a statement that consists of a single expression.
type ExpressionStmt struct {
	Expression Expr
}

func (expr *ExpressionStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitExpressionStmt(expr)
}

// PrintStmt represents a print statement in the AST.
type PrintStmt struct {
	Expression Expr
}

func (expr *PrintStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitPrintStmt(expr)
}

// VarStmt represents a variable declaration statement in the AST.
type VarStmt struct {
	Name        token.Token
	Initializer Expr
}

func (expr *VarStmt) Accept(v ExpressionStmtVisitor) {
	v.VisitVarStmt(expr)
}
