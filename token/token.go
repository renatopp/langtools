package token

import (
	"slices"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func NewToken(t TokenType, l string, line, column int) Token {
	return Token{t, l, line, column}
}

func NewTokenAtChar(t TokenType, l string, char Char) Token {
	return Token{t, l, char.Line, char.Column}
}

func (t *Token) IsOneOfType(tys ...TokenType) bool {
	return slices.Contains(tys, t.Type)
}

func (t *Token) IsOneOfLiterals(lits ...string) bool {
	return slices.Contains(lits, t.Literal)
}

func (t *Token) SetPosition(line, column int) {
	t.Line = line
	t.Column = column
}

func (t *Token) At() (line, column int) {
	return t.Line, t.Column
}
