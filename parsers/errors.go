package parsers

import "github.com/renatopp/langtools/tokens"

type ParserError struct {
	Token tokens.Token
	Msg   string
}

func (e ParserError) Error() string {
	return e.Msg
}

func (e ParserError) At() (line, column int) {
	return e.Token.Line, e.Token.Column
}
