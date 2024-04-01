package lexers

import (
	"errors"
	"unicode/utf8"

	"github.com/renatopp/langtools/tokens"
)

// Scanner is responsible for reading the byte slice input and returning the
// sequence of characters in that input.
type Scanner struct {
	input  []byte
	cursor int
	line   int
	column int
}

// Creates a new scanner attached to the given input.
func NewScanner(input []byte) *Scanner {
	return &Scanner{
		input:  input,
		cursor: 0,
		line:   1,
		column: 1,
	}
}

// Returns true if the cursor is at the end of the input.
func (s *Scanner) IsEof() bool {
	return s.cursor >= len(s.input)
}

// Returns the next character from the input.
func (s *Scanner) Next() (tokens.Char, error) {
	if s.cursor >= len(s.input) {
		return tokens.NewChar(s.line, s.column, 0, 0), nil
	}

	r, size := utf8.DecodeRune(s.input[s.cursor:])
	if r == utf8.RuneError {
		s.cursor += size
		s.column++
		return tokens.NewChar(s.line, s.column, 0, 0), errors.New(ErrInvalidChar)
	}

	c := tokens.NewChar(s.line, s.column, size, r)
	s.cursor += size
	s.column++
	if c.Is('\n') {
		s.line++
		s.column = 1
	}
	return c, nil
}

// Returns all characters from the input. If an error occurs, it will be
// returned instead of the slice.
func (s *Scanner) All() ([]tokens.Char, error) {
	var chars []tokens.Char
	for !s.IsEof() {
		c, err := s.Next()
		if err != nil {
			return nil, err
		}
		chars = append(chars, c)
	}
	return chars, nil
}

// Resets the scanner to the initial state.
func (s *Scanner) Reset() {
	s.cursor = 0
	s.line = 1
	s.column = 1
}
