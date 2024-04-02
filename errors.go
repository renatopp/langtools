package langtools

type Error interface {
	Error() string
	At() (line, column int)
}
