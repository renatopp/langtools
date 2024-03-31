package token

type TokenType string

const (
	TUnknown TokenType = "unknown"
	TEof     TokenType = "eof"
	TString  TokenType = "string"
	TNumber  TokenType = "number"
)
