package langtools

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func NewToken(t TokenType, l string, line, column int) Token {
	return Token{t, l, line, column}
}

type TokenType string

const (
	TUnknown TokenType = "unknown"
	TEof     TokenType = "eof"
	TString  TokenType = "string"
	TNumber  TokenType = "number"
)
