package errors

// Builtin-compatible error interface.
type Error interface {
	Error() string
	At() (line, column int)
	Range() (from, to int)
}
