package parsers

type Parser interface {
	Errors() []ParserError
	HasErrors() bool
}
