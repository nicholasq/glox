package parser

import (
	"github.com/nicholasq/glox/ast"
	err "github.com/nicholasq/glox/error"
	"github.com/nicholasq/glox/token"
)

type Parser struct {
	tokens  []token.Token
	current int
}

func NewParser(tokens *[]token.Token) *Parser {
	return &Parser{
		tokens:  *tokens,
		current: 0,
	}
}

func (p *Parser) Parse() (*[]ast.Stmt, error) {
	var stmts []ast.Stmt
	for !p.isAtEnd() {
		stmts = append(stmts, p.declaration())
	}
	return &stmts, nil
}

func (p *Parser) declaration() ast.Stmt {
	//todo need to catch errors and call p.synchronize in the 'catch' block
	if p.match(token.VAR) {
		return p.varDeclaration()
	}
	return p.statement()
}

func (p *Parser) varDeclaration() ast.Stmt {
	name := p.consume(token.IDENTIFIER, "Expect variable name.")

	var initializer ast.Expr = nil

	if p.match(token.EQUAL) {
		initializer = p.expression()
	}
	p.consume(token.SEMICOLON, "Expect ';' after variable declaration.")
	return &ast.VarStmt{Name: name, Initializer: initializer}
}

func (p *Parser) statement() ast.Stmt {
	if p.match(token.PRINT) {
		stmt := new(ast.PrintStmt)
		*stmt = p.printStatement()
		return stmt
	} else {
		stmt := new(ast.ExpressionStmt)
		*stmt = p.expressionStatement()
		return stmt
	}
}

func (p *Parser) expressionStatement() ast.ExpressionStmt {
	expr := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after expression.")
	return ast.ExpressionStmt{Expression: expr}
}

func (p *Parser) printStatement() ast.PrintStmt {
	expr := p.expression()
	p.consume(token.SEMICOLON, "Expect ';' after value.")
	return ast.PrintStmt{Expression: expr}
}

/*
	Grammar:
	program        -> declaration* EOF ;
    declaration    -> varDecl | statement;
	varDecl 	   -> "var" IDENTIFIER ( "=" expression )? ";" ;
	statement      -> exprStmt | printStmt;
	expression     → equality ;
	equality       → comparison ( ( "!=" | "==" ) comparison )* ;
	comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
	term           → factor ( ( "-" | "+" ) factor )* ;
	factor         → unary ( ( "/" | "*" ) unary )* ;
	unary          → ( "!" | "-" ) unary
				   | primary ;
	primary        → NUMBER | STRING | "true" | "false" | "nil"
				   | "(" expression ")" | IDENTIFIER ;
*/

// expression parses and returns an expression.
// It calls the equality method to handle equality operators.
// If there are multiple equality operators, it iterates over them and
// constructs a Binary expression.
// Returns the parsed expression.
func (p *Parser) expression() ast.Expr {
	return p.equality()
}

// equality parses and returns an expression.
// It calls the comparison method to handle comparison operators.
// If there are multiple equality operators, it iterates over them and
// constructs a Binary expression.
// Returns the parsed expression.
func (p *Parser) equality() ast.Expr {
	expr := p.comparison()
	for p.match(token.BANG_EQUAL, token.EQUAL_EQUAL) {
		operator := p.previous()
		right := p.comparison()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: right}
	}
	return expr
}

// comparison parses and returns an expression.
// It calls the term method to handle term operators.
// If there are multiple term operators, it iterates over them and
// constructs a Binary expression.
// Returns the parsed expression.
func (p *Parser) comparison() ast.Expr {
	expr := p.term()
	for p.match(token.GREATER, token.GREATER_EQUAL, token.LESS, token.LESS_EQUAL) {
		operator := p.previous()
		term := p.term()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: term}
	}
	return expr
}

// term parses and returns a term of an expression.
// It calls the factor method to handle unary operators and operands.
// If there are multiple term operators, it iterates over them and constructs a Binary expression.
// Returns the parsed term expression.
func (p *Parser) term() ast.Expr {
	expr := p.factor()
	for p.match(token.MINUS, token.PLUS) {
		operator := p.previous()
		factor := p.factor()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: factor}
	}
	return expr
}

// factor parses and returns a factor of an expression.
// It calls the unary method to handle unary operators and operands.
// If there are multiple factor operators, it iterates over them and constructs a Binary expression.
// Returns the parsed factor expression.
func (p *Parser) factor() ast.Expr {
	expr := p.unary()
	for p.match(token.SLASH, token.STAR) {
		operator := p.previous()
		factor := p.unary()
		expr = &ast.Binary{Left: expr, Operator: operator, Right: factor}
	}
	return expr
}

// unary parses and returns a unary expression.
// If the current token is a BANG or MINUS token, it consumes the token,
// recursively calls unary to parse the operand, and returns a Unary expression.
// Otherwise, it calls primary method to parse the primary expression.
// Returns the parsed unary expression.
func (p *Parser) unary() ast.Expr {
	if p.match(token.BANG, token.MINUS) {
		operator := p.previous()
		right := p.unary()
		return &ast.Unary{Operator: operator, Right: right}
	}
	return p.primary()
}

func (p *Parser) primary() ast.Expr {
	if p.match(token.FALSE) {
		return &ast.Literal{Value: false}
	}
	if p.match(token.TRUE) {
		return &ast.Literal{Value: true}
	}
	if p.match(token.NIL) {
		return &ast.Literal{Value: nil}
	}
	if p.match(token.NUMBER, token.STRING) {
		return &ast.Literal{Value: p.previous().Literal}
	}
	if p.match(token.IDENTIFIER) {
		return &ast.Variable{Name: p.previous()}
	}
	if p.match(token.LEFT_PAREN) {
		expr := p.expression()
		p.consume(token.RIGHT_PAREN, "Expect ')' after expression.")
		return &ast.Grouping{Expression: expr}
	}

	p.logError(p.peek(), "Expect expression.")
	panic("We shouldn't have gotten here...")
}

func (p *Parser) logError(token token.Token, message string) {
	//todo call glox.error
	err.GloxError(token, message)
	panic("Parse error")
}

func (p *Parser) consume(tokenType token.TokenType, message string) token.Token {
	if p.check(tokenType) {
		return p.advance()
	}
	p.logError(p.peek(), message)
	panic("Parse error")
}

func (p *Parser) match(tokenType ...token.TokenType) bool {
	for _, token := range tokenType {
		if p.check(token) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType token.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().TokenType == tokenType
}

func (p *Parser) advance() token.Token {
	p.current++
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().TokenType == token.EOF
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) synchronize() {
	p.advance()
	for !p.isAtEnd() {
		if p.previous().TokenType == token.SEMICOLON {
			return
		}
		switch p.peek().TokenType {
		case token.CLASS:
		case token.FUN:
		case token.VAR:
		case token.FOR:
		case token.IF:
		case token.WHILE:
		case token.PRINT:
		case token.RETURN:
			return
		}
		p.advance()
	}
}
