package tokens

import "slices"

// Represents a character in the source code, with its position.
type Char struct {
	Rune   rune
	Size   int // Size in bytes
	Line   int
	Column int
}

// Creates a new Char at line and column with the given rune and size.
func NewChar(line, column, size int, rune rune) Char {
	return Char{
		Rune:   rune,
		Size:   size,
		Line:   line,
		Column: column,
	}
}

// Returns the line and column of the character.
func (p *Char) At() (line, column int) {
	return p.Line, p.Column
}

// Checks if the character rune matches the given rune.
func (p *Char) Is(r rune) bool {
	return p.Rune == r
}

// Checks if the character rune is one of the given runes.
func (p *Char) IsOneOf(runes ...rune) bool {
	return slices.Contains(runes, p.Rune)
}
