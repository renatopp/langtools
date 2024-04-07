package errors

// Builtin-compatible error interface.
type Error interface {
	Error() string                                       // Returns the error message.
	At() (line, column int)                              // Returns the initial position of the error.
	Range() (fromLine, fromColumn, toLine, toColumn int) // Returns the range of the error.
}
