package parser

import (
	"reflect"
	"testing"

	"github.com/nicholasq/glox/ast"
	"github.com/nicholasq/glox/scanner"
	"github.com/nicholasq/glox/token"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []ast.Stmt
	}{
		{
			name:  "Variable declaration",
			input: "var min = 1;",
			expected: []ast.Stmt{
				&ast.VarStmt{
					Name:        token.Token{TokenType: token.IDENTIFIER, Lexeme: "min", Literal: "min", Line: 1},
					Initializer: &ast.Literal{Value: float64(1)},
				},
			},
		},
		{
			name:  "Variable declaration with binary expression",
			input: "var add = 1 + 5;",
			expected: []ast.Stmt{
				&ast.VarStmt{
					Name: token.Token{TokenType: token.IDENTIFIER, Lexeme: "add", Literal: "add", Line: 1},
					Initializer: &ast.Binary{
						Left:     &ast.Literal{Value: float64(1)},
						Operator: token.Token{TokenType: token.PLUS, Lexeme: "+", Literal: "+", Line: 1},
						Right:    &ast.Literal{Value: float64(5)},
					},
				},
			},
		},
		{
			name:  "Variable declaration with precedence and grouping",
			input: "var average = (min + max) / 2;",
			expected: []ast.Stmt{
				&ast.VarStmt{
					Name: token.Token{TokenType: token.IDENTIFIER, Lexeme: "average", Literal: "average", Line: 1},
					Initializer: &ast.Binary{
						Left: &ast.Grouping{
							Expression: &ast.Binary{
								Left:     &ast.Variable{Name: token.Token{TokenType: token.IDENTIFIER, Lexeme: "min", Literal: "min", Line: 1}},
								Operator: token.Token{TokenType: token.PLUS, Lexeme: "+", Literal: "+", Line: 1},
								Right:    &ast.Variable{Name: token.Token{TokenType: token.IDENTIFIER, Lexeme: "max", Literal: "max", Line: 1}},
							},
						},
						Operator: token.Token{TokenType: token.SLASH, Lexeme: "/", Literal: "/", Line: 1},
						Right:    &ast.Literal{Value: float64(2)},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := scanner.New(tt.input)
			tokens := scanner.ScanTokens()
			parser := New(tokens)
			result, err := parser.Parse()

			if err != nil {
				t.Fatalf("Error during parsing: %s", err)
			}

			compareAST(t, tt.expected, result)
		})
	}
}

func compareAST(t *testing.T, expected, actual []ast.Stmt) {
	if len(expected) != len(actual) {
		t.Fatalf("Expected %d statements, got %d", len(expected), len(actual))
	}

	for i := range expected {
		compareNode(t, expected[i], actual[i])
	}
}

func compareNode(t *testing.T, expected, actual interface{}) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		t.Fatalf("Expected type %T, got %T", expected, actual)
	}

	switch exp := expected.(type) {
	case *ast.VarStmt:
		act := actual.(*ast.VarStmt)
		compareVarStmt(t, exp, act)
	case *ast.ExpressionStmt:
		act := actual.(*ast.ExpressionStmt)
		compareExpressionStmt(t, exp, act)
	case *ast.PrintStmt:
		act := actual.(*ast.PrintStmt)
		comparePrintStmt(t, exp, act)
	case *ast.Literal:
		act := actual.(*ast.Literal)
		compareLiteral(t, exp, act)
	case *ast.Binary:
		act := actual.(*ast.Binary)
		compareBinary(t, exp, act)
	case *ast.Grouping:
		act := actual.(*ast.Grouping)
		compareGrouping(t, exp, act)
	case *ast.Variable:
		act := actual.(*ast.Variable)
		compareVariable(t, exp, act)
	default:
		t.Fatalf("Unsupported node type: %T", exp)
	}
}

func compareVarStmt(t *testing.T, expected, actual *ast.VarStmt) {
	if expected.Name != actual.Name {
		t.Fatalf("Expected name %v, got %v", expected.Name, actual.Name)
	}
	if expected.Initializer == nil && actual.Initializer == nil {
		return
	}
	if expected.Initializer == nil || actual.Initializer == nil {
		t.Fatalf("Initializer mismatch: expected %v, got %v", expected.Initializer, actual.Initializer)
	}
	compareNode(t, expected.Initializer, actual.Initializer)
}

func compareExpressionStmt(t *testing.T, expected, actual *ast.ExpressionStmt) {
	compareNode(t, expected.Expression, actual.Expression)
}

func comparePrintStmt(t *testing.T, expected, actual *ast.PrintStmt) {
	compareNode(t, expected.Expression, actual.Expression)
}

func compareLiteral(t *testing.T, expected, actual *ast.Literal) {
	if !reflect.DeepEqual(expected.Value, actual.Value) {
		t.Fatalf("Expected literal value %+v, got %+v", expected.Value, actual.Value)
	}
}

func compareBinary(t *testing.T, expected, actual *ast.Binary) {
	compareNode(t, expected.Left, actual.Left)
	if expected.Operator != actual.Operator {
		t.Fatalf("Expected operator %v, got %v", expected.Operator, actual.Operator)
	}
	compareNode(t, expected.Right, actual.Right)
}

func compareGrouping(t *testing.T, expected, actual *ast.Grouping) {
	if expected.Expression == nil && actual.Expression == nil {
		return
	}
	if expected.Expression == nil || actual.Expression == nil {
		t.Fatalf("Grouping expression mismatch: expected %v, got %v", expected.Expression, actual.Expression)
	}
	compareNode(t, expected.Expression, actual.Expression)
}

func compareVariable(t *testing.T, expected, actual *ast.Variable) {
	if expected.Name != actual.Name {
		t.Fatalf("Expected variable name %v, got %v", expected.Name, actual.Name)
	}
}

// Add more comparison functions for other expression types (Binary, Unary, etc.)
