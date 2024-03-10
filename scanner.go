package main

type Scanner struct {
	Source  string
	Tokens  []Token
	Start   uint
	Current uint
	Line    uint
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

func MakeScanner(source string) Scanner {
	return Scanner{
		Source:  source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	// Begin scanning the Source at most 2 characters at a time.
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.Start = s.Current
		s.scanToken()
	}
	// We are done scanning the Source.
	// Apply the EOF token.
	s.Tokens = append(s.Tokens, Token{TokenType: EOF, Lexeme: "", Literal: nil, Line: s.Line})
	return s.Tokens
}

func (s *Scanner) scanToken() {
	/*
		We Start by ingesting lexemes.
		We check to see if the Current character is either
		a (single/multi-)character (operator or delimiter) lexeme.
		If it is none of the recognized operators or delimiter lexemes,
		we check to see if it is the beginning of a literal or keyword lexeme.
	*/
	c := s.getRuneAndAdvance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the Line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.getRuneAndAdvance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		break
	case '\n':
		s.Line++
		break
	case '"':
		s.string()
		break
	default:
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			report(s.Line, "", "Unexpected character.")
		}
	}
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.getRuneAndAdvance()
	}
	text := s.Source[s.Start:s.Current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}
	s.addToken(tokenType)
}

func (s *Scanner) number() {
	for s.isDigit(s.peek()) {
		s.getRuneAndAdvance()
	}
	// Look for a fractional part.
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.getRuneAndAdvance()
		for s.isDigit(s.peek()) {
			s.getRuneAndAdvance()
		}
	}
	s.addTokenLiteral(NUMBER, s.Source[s.Start:s.Current])
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.Line++
		}
		s.getRuneAndAdvance()
	}
	// Unterminated string.
	if s.isAtEnd() {
		report(s.Line, "", "Unterminated string.")
		return
	}
	// The closing ".
	s.getRuneAndAdvance()
	// Trim the surrounding quotes.
	value := s.Source[s.Start+1 : s.Current-1]
	s.addTokenLiteral(STRING, value)
}

func (s *Scanner) match(char rune) bool {
	if s.isAtEnd() {
		return false
	}
	if []rune(s.Source)[s.Current] != char {
		return false
	}
	s.Current++
	return true
}

func (s *Scanner) peek() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return []rune(s.Source)[s.Current]
}

func (s *Scanner) peekNext() rune {
	if s.isAtEnd() {
		return '\000'
	}
	return []rune(s.Source)[s.Current+1]
}

func (s *Scanner) isAlpha(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func (s *Scanner) isAlphaNumeric(char rune) bool {
	return s.isAlpha(char) || s.isDigit(char)
}

func (s *Scanner) isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func (s *Scanner) isAtEnd() bool { return s.Current >= uint(len(s.Source)) }

func (s *Scanner) getRuneAndAdvance() rune {
	curr := []rune(s.Source)[s.Current]
	s.Current++
	return curr
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: s.Line})
}
