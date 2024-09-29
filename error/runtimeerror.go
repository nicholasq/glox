package error

import "github.com/nicholasq/glox/token"

type RuntimeError struct {
	Token   token.Token
	Message string
}
