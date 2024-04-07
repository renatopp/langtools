package lexers

import (
	"github.com/renatopp/langtools/tokens"
)

type Lexer interface {
	Errors() []LexerError
	HasErrors() bool
	Next() (*tokens.Token, bool)
	EatToken() *tokens.Token
	PeekToken() *tokens.Token
	PeekTokenAt(int) *tokens.Token
}
