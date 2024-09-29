package environment

import (
	"errors"
	"fmt"

	"github.com/nicholasq/glox/token"
)

type Environment struct {
	values map[string]interface{}
}

func (e *Environment) Define(name string, value interface{}) {
	e.values[name] = value
}

func (e *Environment) Get(name token.Token) (interface{}, error) {
	if val, ok := e.values[name.Lexeme]; ok {
		return val, nil
	}
	return nil, errors.New(fmt.Sprintf("undefined variable: %v", name.Lexeme))
}
