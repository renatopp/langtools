package tokens

import (
	"fmt"
	"slices"
)

// Represents a token in the source code.
//
//	tokens.NewToken(EOF, "").WithRangeChars(from, to)
type Token struct {
	Type       TokenType
	Literal    string
	FromLine   int
	FromColumn int
	ToLine     int
	ToColumn   int
}

// Creates a new token.
func NewToken(t TokenType, l string) *Token {
	return &Token{
		Type:    t,
		Literal: l,
	}
}

// Creates a new token with the given range. You can use this method as a chain.
func (t Token) WithRange(fromLine, fromColumn, toLine, toColumn int) *Token {
	t.FromLine = fromLine
	t.FromColumn = fromColumn
	t.ToLine = toLine
	t.ToColumn = toColumn
	return &t
}

// Creates a new token with the given char range. You can use this method as a
// chain.
func (t Token) WithRangeChars(from, to Char) *Token {
	t.FromLine = from.Line
	t.FromColumn = from.Column
	t.ToLine = to.Line
	t.ToColumn = to.Column
	return &t
}

// Creates a new token with the given type. You can use this method as a chain.
func (t Token) WithType(tp TokenType) *Token {
	t.Type = tp
	return &t
}

// Creates a new token with the given literal. You can use this method as a
// chain.
func (t Token) WithLiteral(lit string) *Token {
	t.Literal = lit
	return &t
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
func (t *Token) IsOneOfTypes(tys ...TokenType) bool {
	return slices.Contains(tys, t.Type)
}

// Returns true if the token is one of the given literals.
func (t *Token) IsOneOfLiterals(lits ...string) bool {
	return slices.Contains(lits, t.Literal)
}

// Returns the initial position of the token.
func (t *Token) From() (line, column int) {
	return t.FromLine, t.FromColumn
}

// Returns the final position of the token.
func (t *Token) To() (line, column int) {
	return t.ToLine, t.ToColumn
}

// Returns the range of the starting and final lines. Starting at 1.
func (t *Token) RangeLines() (from, to int) {
	return t.FromLine, t.ToLine
}

// Returns the range of the token ignoring the lines. Starting at 1.
func (t *Token) RangeColumns() (from, to int) {
	return t.FromColumn, t.ToColumn
}

// Returns the range of the token.
func (t *Token) Range() (fromLine, fromColumn, toLine, toColumn int) {
	return t.FromLine, t.FromColumn, t.ToLine, t.ToColumn
}

// Pretty string representation of the token.
func (t *Token) DebugString() string {
	return fmt.Sprintf("[%d:%d] (%s) %s", t.FromLine, t.FromColumn, t.Type, t.Literal)
}
