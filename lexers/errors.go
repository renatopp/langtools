package lexers

var (
	ErrInvalidChar         = "invalid UTF-8 encoding"
	ErrUnexpectedNewline   = "unexpected newline"
	ErrUnexpectedEndOfFile = "unexpected end of file"
	ErrUnexpectedDot       = "unexpected '.' character"
	ErrUnexpectedE         = "unexpected 'e' character"
)

type LexerError struct {
	Line   int
	Column int
	Msg    string
}

func NewLexerError(line, column int, msg string) LexerError {
	return LexerError{
		Line:   line,
		Column: column,
		Msg:    msg,
	}
}

func (e LexerError) Error() string {
	return e.Msg
}

func (e LexerError) At() (line, column int) {
	return e.Line, e.Column
}

func (e LexerError) Range() (fromLine, fromColumn, toLine, toColumn int) {
	return e.Line, e.Column, e.Line, e.Column + 1
}
