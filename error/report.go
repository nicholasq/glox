package error

import (
	"fmt"

	"github.com/nicholasq/glox/token"
)

func ErrorReport(Line uint, message string) {
	//todo: implement
	//report(Line, "", message)
}

func Report(line uint, where string, message string) bool {
	fmt.Printf("[Line %d] Error %s: %s\n", line, where, message)
	return true
}

func GloxError(tok token.Token, message string) {
	if tok.TokenType == token.EOF {
		Report(tok.Line, " at end", message)
	} else {
		Report(tok.Line, " at '"+tok.Lexeme+"'", message)
	}
}
