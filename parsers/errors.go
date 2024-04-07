package parsers

import (
	"github.com/renatopp/langtools/tokens"
)

type ParserError struct {
	Token *tokens.Token
	Msg   string
}

// Creates a new parser error with the given message.
func NewParserError(token *tokens.Token, msg string) ParserError {
	return ParserError{
		Token: token,
		Msg:   msg,
	}
}

// Returns the error message.
func (e ParserError) Error() string {
	return e.Msg
}

// Returns the initial position of the error.
func (e ParserError) At() (line, column int) {
	return e.Token.From()
}

// Returns the range of the error.
func (e ParserError) Range() (fromLine, fromColumn, toLine, toColumn int) {
	return e.Token.Range()
}
