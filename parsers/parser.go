package parsers

import (
	"github.com/renatopp/langtools/lexers"
)

type BaseParser struct {
	MaxErrors int
	lexer     lexers.Lexer
	errors    []ParserError
}

// Creates a new parser with the given lexer.
func NewBaseParser(lexer lexers.Lexer) *BaseParser {
	return &BaseParser{
		MaxErrors: 10,
		lexer:     lexer,
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
