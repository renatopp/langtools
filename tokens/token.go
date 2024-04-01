package tokens

import (
	"fmt"
	"slices"
)

// Represents a token in the source code.
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Creates a new token.
func NewToken(t TokenType, l string, line, column int) Token {
	return Token{t, l, line, column}
}

// Creates a new token at the same position as the given character.
func NewTokenAtChar(t TokenType, l string, char Char) Token {
	return Token{t, l, char.Line, char.Column}
}

// Creates a new EOF token.
func NewEofToken() Token {
	return Token{EOF, "", 0, 0}
}

// Creates a new EOF token at the same position as the given character.
func NewEofTokenAtChar(char Char) Token {
	return Token{EOF, "", char.Line, char.Column}
}

// Returns true if the token is of the given type.
func (t *Token) IsType(ty TokenType) bool {
	return t.Type == ty
}

// Returns true if the token is the given literal.
func (t *Token) IsLiteral(lit string) bool {
	return t.Literal == lit
}

// Returns true if the token is of the given type.
func (t *Token) IsOneOfType(tys ...TokenType) bool {
	return slices.Contains(tys, t.Type)
}

// Returns true if the token is one of the given literals.
func (t *Token) IsOneOfLiterals(lits ...string) bool {
	return slices.Contains(lits, t.Literal)
}

// Set the position of the token.
func (t *Token) SetPosition(line, column int) {
	t.Line = line
	t.Column = column
}

// Returns the position of the token.
func (t *Token) At() (line, column int) {
	return t.Line, t.Column
}

// Converts the token to a given type.
func (t Token) As(tp TokenType) Token {
	t.Type = tp
	return t
}

// Pretty string representation of the token.
func (t Token) DebugString() string {
	return fmt.Sprintf("[%d:%d] (%s) %s", t.Line, t.Column, t.Type, t.Literal)
}
