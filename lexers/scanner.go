package lexers

import (
	"errors"
	"unicode/utf8"

	"github.com/renatopp/langtools/tokens"
)

type Scanner struct {
	input  []byte
	cursor int
	line   int
	column int
}

func NewScanner(input []byte) *Scanner {
	return &Scanner{
		input:  input,
		cursor: 0,
		line:   1,
		column: 1,
	}
}

// Returns true if the cursor is at the end of the input.
func (l *Scanner) IsEof() bool {
	return l.cursor >= len(l.input)
}

// Returns the next valid token from the input. If the input is empty, or the
// Scanner has too many errors, or the file ended it returns an empty char.
func (l *Scanner) Next() (tokens.Char, error) {
	if l.cursor >= len(l.input) {
		return tokens.NewChar(l.line, l.column, 0, 0), nil
	}

	r, size := utf8.DecodeRune(l.input[l.cursor:])
	if r == utf8.RuneError {
		l.cursor += size
		l.column++
		return tokens.NewChar(l.line, l.column, 0, 0), errors.New(ErrInvalidChar)
	}

	c := tokens.NewChar(l.line, l.column, size, r)
	l.cursor += size
	l.column++
	if c.Is('\n') {
		l.line++
		l.column = 1
	}
	return c, nil
}
