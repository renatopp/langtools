package token

// Represents a character in the source code, with its position.
type Char struct {
	Rune   rune
	Size   int // Size in bytes
	Line   int
	Column int
}

func NewChar(line, column, size int, rune rune) Char {
	return Char{
		Rune:   rune,
		Size:   size,
		Line:   line,
		Column: column,
	}
}

func (p *Char) At() (line, column int) {
	return p.Line, p.Column
}

func (p *Char) Is(r rune) bool {
	return p.Rune == r
}
