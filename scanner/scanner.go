package scanner

import (
	"strconv"

	"github.com/nicholasq/glox/error"
	"github.com/nicholasq/glox/token"
)

type Scanner struct {
	Source  string
	Tokens  []token.Token
	Start   uint
	Current uint
	Line    uint
}

func New(source string) Scanner {
	return Scanner{
		Source:  source,
		Tokens:  []token.Token{},
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	// Begin scanning the Source at most 2 characters at a time.
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.Start = s.Current
		s.scanToken()
	}
	// We are done scanning the Source.
	// Apply the EOF token.
	s.Tokens = append(s.Tokens, token.Token{TokenType: token.EOF, Lexeme: "", Literal: nil, Line: s.Line})
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
		s.addToken(token.LEFT_PAREN)
	case ')':
		s.addToken(token.RIGHT_PAREN)
	case '{':
		s.addToken(token.LEFT_BRACE)
	case '}':
		s.addToken(token.RIGHT_BRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.nextRuneMatches('=') {
			s.addToken(token.BANG_EQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.nextRuneMatches('=') {
			s.addToken(token.EQUAL_EQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.nextRuneMatches('=') {
			s.addToken(token.LESS_EQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.nextRuneMatches('=') {
			s.addToken(token.GREATER_EQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case '/':
		if s.nextRuneMatches('/') {
			// A comment goes until the end of the Line.
			for s.peek() != '\n' && !s.isAtEnd() {
				s.getRuneAndAdvance()
			}
		} else {
			s.addToken(token.SLASH)
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
			error.Report(s.Line, "", "Unexpected character.")
		}
	}
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.getRuneAndAdvance()
	}
	text := s.Source[s.Start:s.Current]
	tokenType, ok := token.Keywords[text]
	if !ok {
		tokenType = token.IDENTIFIER
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
	str := s.Source[s.Start:s.Current]
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err) // todo add better error handling
	}
	s.addTokenLiteral(token.NUMBER, f)
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
		error.Report(s.Line, "", "Unterminated string.")
		return
	}
	// Consumes the closing '"'.
	s.getRuneAndAdvance()
	// Trim the surrounding quotes.
	value := s.Source[s.Start+1 : s.Current-1]
	s.addTokenLiteral(token.STRING, value)
}

func (s *Scanner) nextRuneMatches(char rune) bool {
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

func (s *Scanner) addToken(tokenType token.TokenType) {
	literal := s.Source[s.Start:s.Current]
	s.addTokenLiteral(tokenType, literal)
}

func (s *Scanner) addTokenLiteral(tokenType token.TokenType, literal interface{}) {
	text := s.Source[s.Start:s.Current]
	s.Tokens = append(s.Tokens, token.Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: s.Line})
}
