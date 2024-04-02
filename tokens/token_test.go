package tokens_test

import (
	"testing"

	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "", 1, 2)
	assert.Equal(t, tokens.Token{tokens.EOF, "", 1, 2}, token)
}

func TestNewTokenAtChar(t *testing.T) {
	char := tokens.NewChar(1, 2, 3, 'a')
	token := tokens.NewTokenAtChar(tokens.EOF, "", char)
	assert.Equal(t, tokens.Token{tokens.EOF, "", 1, 2}, token)
}

func TestNewEofToken(t *testing.T) {
	token := tokens.NewEofToken()
	assert.Equal(t, tokens.Token{tokens.EOF, "", 0, 0}, token)
}

func TestNewEofTokenAtChar(t *testing.T) {
	char := tokens.NewChar(1, 2, 3, 'a')
	token := tokens.NewEofTokenAtChar(char)
	assert.Equal(t, tokens.Token{tokens.EOF, "", 1, 2}, token)
}

func TestIsType(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "", 1, 2)
	assert.True(t, token.IsType(tokens.EOF))
}

func TestIsLiteral(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "renato", 1, 2)
	assert.True(t, token.IsLiteral("renato"))
}

func TestTokenIsOneOfType(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "", 1, 2)
	assert.True(t, token.IsOneOfType(tokens.EOF))
}

func TestTokenIsOneOfLiterals(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "renato", 1, 2)
	assert.True(t, token.IsOneOfLiterals("sample", "renato"))
}

func TestAt(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "", 1, 2)
	line, column := token.At()
	assert.Equal(t, 1, line)
	assert.Equal(t, 2, column)
}

func TestAs(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "", 1, 2)
	assert.Equal(t, tokens.UNKNOWN, token.As(tokens.UNKNOWN).Type)
}
