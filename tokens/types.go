package tokens

type TokenType string

const (
	UNKNOWN TokenType = "unknown"
	INVALID TokenType = "invalid"
	EOF     TokenType = "eof"
)
