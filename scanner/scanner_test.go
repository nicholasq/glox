package scanner

import (
	"strconv"
	"testing"

	"github.com/nicholasq/glox/token"
)

func TestScanTokens(t *testing.T) {
	source := `var one = 1;
	var bool = true;

	if (bool) {
		print "yes";
	} else {
		print "no";
	}

	while (bool) {
		print 345;
	}

	fun double(num) {
	 	return num * 2;
	}

	double(one);
	`

	var toFloat = func(literal string) float64 {
		oneVal, err := strconv.ParseFloat(literal, 64)
		if err != nil {
			t.Fatalf("could not convert '%s' to a float", literal)
		}
		return oneVal
	}

	scanner := New(source)
	tokens := scanner.ScanTokens()
	expected := []token.Token{
		{
			TokenType: token.VAR,
			Lexeme:    "var",
			Literal:   "var",
			Line:      1,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "one",
			Literal:   "one",
			Line:      1,
		},
		{
			TokenType: token.EQUAL,
			Lexeme:    "=",
			Literal:   "=",
			Line:      1,
		},
		{
			TokenType: token.NUMBER,
			Lexeme:    "1",
			Literal:   toFloat("1"),
			Line:      1,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      1,
		},
		{
			TokenType: token.VAR,
			Lexeme:    "var",
			Literal:   "var",
			Line:      2,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "bool",
			Literal:   "bool",
			Line:      2,
		},
		{
			TokenType: token.EQUAL,
			Lexeme:    "=",
			Literal:   "=",
			Line:      2,
		},
		{
			TokenType: token.TRUE,
			Lexeme:    "true",
			Literal:   "true",
			Line:      2,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      2,
		},
		{
			TokenType: token.IF,
			Lexeme:    "if",
			Literal:   "if",
			Line:      4,
		},
		{
			TokenType: token.LEFT_PAREN,
			Lexeme:    "(",
			Literal:   "(",
			Line:      4,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "bool",
			Literal:   "bool",
			Line:      4,
		},
		{
			TokenType: token.RIGHT_PAREN,
			Lexeme:    ")",
			Literal:   ")",
			Line:      4,
		},
		{
			TokenType: token.LEFT_BRACE,
			Lexeme:    "{",
			Literal:   "{",
			Line:      4,
		},
		{
			TokenType: token.PRINT,
			Lexeme:    "print",
			Literal:   "print",
			Line:      5,
		},
		{
			TokenType: token.STRING,
			Lexeme:    "\"yes\"",
			Literal:   "yes",
			Line:      5,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      5,
		},
		{
			TokenType: token.RIGHT_BRACE,
			Lexeme:    "}",
			Literal:   "}",
			Line:      6,
		},
		{
			TokenType: token.ELSE,
			Lexeme:    "else",
			Literal:   "else",
			Line:      6,
		},
		{
			TokenType: token.LEFT_BRACE,
			Lexeme:    "{",
			Literal:   "{",
			Line:      6,
		},
		{
			TokenType: token.PRINT,
			Lexeme:    "print",
			Literal:   "print",
			Line:      7,
		},
		{
			TokenType: token.STRING,
			Lexeme:    "\"no\"",
			Literal:   "no",
			Line:      7,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      7,
		},
		{
			TokenType: token.RIGHT_BRACE,
			Lexeme:    "}",
			Literal:   "}",
			Line:      8,
		},
		{
			TokenType: token.WHILE,
			Lexeme:    "while",
			Literal:   "while",
			Line:      10,
		},
		{
			TokenType: token.LEFT_PAREN,
			Lexeme:    "(",
			Literal:   "(",
			Line:      10,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "bool",
			Literal:   "bool",
			Line:      10,
		},
		{
			TokenType: token.RIGHT_PAREN,
			Lexeme:    ")",
			Literal:   ")",
			Line:      10,
		},
		{
			TokenType: token.LEFT_BRACE,
			Lexeme:    "{",
			Literal:   "{",
			Line:      10,
		},
		{
			TokenType: token.PRINT,
			Lexeme:    "print",
			Literal:   "print",
			Line:      11,
		},
		{
			TokenType: token.NUMBER,
			Lexeme:    "345",
			Literal:   toFloat("345"),
			Line:      11,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      11,
		},
		{
			TokenType: token.RIGHT_BRACE,
			Lexeme:    "}",
			Literal:   "}",
			Line:      12,
		},
		{
			TokenType: token.FUN,
			Lexeme:    "fun",
			Literal:   "fun",
			Line:      14,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "double",
			Literal:   "double",
			Line:      14,
		},
		{
			TokenType: token.LEFT_PAREN,
			Lexeme:    "(",
			Literal:   "(",
			Line:      14,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "num",
			Literal:   "num",
			Line:      14,
		},
		{
			TokenType: token.RIGHT_PAREN,
			Lexeme:    ")",
			Literal:   ")",
			Line:      14,
		},
		{
			TokenType: token.LEFT_BRACE,
			Lexeme:    "{",
			Literal:   "{",
			Line:      14,
		},
		{
			TokenType: token.RETURN,
			Lexeme:    "return",
			Literal:   "return",
			Line:      15,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "num",
			Literal:   "num",
			Line:      15,
		},
		{
			TokenType: token.STAR,
			Lexeme:    "*",
			Literal:   "*",
			Line:      15,
		},
		{
			TokenType: token.NUMBER,
			Lexeme:    "2",
			Literal:   toFloat("2"),
			Line:      15,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      15,
		},
		{
			TokenType: token.RIGHT_BRACE,
			Lexeme:    "}",
			Literal:   "}",
			Line:      16,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "double",
			Literal:   "double",
			Line:      18,
		},
		{
			TokenType: token.LEFT_PAREN,
			Lexeme:    "(",
			Literal:   "(",
			Line:      18,
		},
		{
			TokenType: token.IDENTIFIER,
			Lexeme:    "one",
			Literal:   "one",
			Line:      18,
		},
		{
			TokenType: token.RIGHT_PAREN,
			Lexeme:    ")",
			Literal:   ")",
			Line:      18,
		},
		{
			TokenType: token.SEMICOLON,
			Lexeme:    ";",
			Literal:   ";",
			Line:      18,
		},
		{
			TokenType: token.EOF,
			Lexeme:    "",
			Literal:   nil,
			Line:      19,
		},
	}

	if len(tokens) != len(expected) {
		t.Fatalf("Expected %d tokens. Got %d", len(tokens), len(expected))
	}

	for idx, token := range tokens {
		if !deepEqual(expected[idx], token) {
			t.Fatalf("\nExpected: %+v\n     Got: %+v", expected[idx], token)
		}
	}
}

func deepEqual(expected, actual token.Token) bool {
	if expected.TokenType != actual.TokenType {
		return false
	}

	if expected.Literal != actual.Literal {
		return false
	}

	if expected.Lexeme != actual.Lexeme {
		return false
	}

	if expected.Line != actual.Line {
		return false
	}

	return true
}
