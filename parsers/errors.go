package parsers

import "github.com/renatopp/langtools/token"

type ParserError struct {
	Token token.Token
	Msg   string
}

func (e *ParserError) Error() string {
	return e.Msg
}

func (e *ParserError) At() (line, column int) {
	return e.Token.Line, e.Token.Column
}
