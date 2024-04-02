package parsers

import (
	"fmt"
	"strings"

	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/tokens"
)

type BaseParser struct {
	MaxErrors int
	Lexer     lexers.Lexer
	errors    []ParserError
}

// Creates a new parser with the given lexer.
func NewBaseParser(lexer lexers.Lexer) *BaseParser {
	return &BaseParser{
		MaxErrors: 10,
		Lexer:     lexer,
		errors:    make([]ParserError, 0),
	}
}

// Returns the errors found by the lexer.
func (p *BaseParser) Errors() []ParserError {
	return p.errors
}

// Returns true if the lexer has errors.
func (p *BaseParser) HasErrors() bool {
	return len(p.errors) > 0
}

// Returns true if the lexer has too many errors.
func (p *BaseParser) HasTooManyErrors() bool {
	return len(p.errors) >= p.MaxErrors
}

// Register an error with the given message.
func (p *BaseParser) RegisterError(msg string) {
	p.errors = append(p.errors, ParserError{
		Token: p.Lexer.PeekToken(),
		Msg:   msg,
	})
}

// Register an error with the given message and token.
func (p *BaseParser) RegisterErrorWithToken(msg string, token tokens.Token) {
	p.errors = append(p.errors, ParserError{
		Token: token,
		Msg:   msg,
	})
}

// Checks if the next token is the given types.
func (p *BaseParser) ExpectType(expected ...tokens.TokenType) bool {
	cur := p.Lexer.PeekToken()
	for _, t := range expected {
		if cur.IsType(t) {
			return true
		}
	}

	e := make([]string, len(expected))
	for i, t := range expected {
		e[i] = string(t)
	}

	p.RegisterErrorWithToken(fmt.Sprintf("expected one of the following tokens: [%s]", strings.Join(e, ", ")), cur)
	return false
}
