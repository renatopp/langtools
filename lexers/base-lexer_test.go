package lexers_test

import (
	"testing"

	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

var helloInput = []byte("Hello, World!")
var helloOutput = []rune("Hello, World!")

func TestErrorRegistration(t *testing.T) {
	lexer := lexers.NewBaseLexer(helloInput, nil)
	lexer.MaxErrors = 2

	assert.False(t, lexer.HasErrors())
	assert.False(t, lexer.HasTooManyErrors())

	lexer.RegisterError("Error 1")
	lexer.RegisterError("Error 2")
	lexer.RegisterError("Error 3")

	assert.True(t, lexer.HasErrors())
	assert.True(t, lexer.HasTooManyErrors())
	assert.Len(t, lexer.Errors(), 2)
}

func TestReadChars(t *testing.T) {
	lexer := lexers.NewBaseLexer(helloInput, nil)

	assert.Equal(t, helloOutput[0], lexer.PeekChar().Rune)
	assert.Equal(t, helloOutput[0], lexer.PeekChar().Rune)
	assert.Equal(t, helloOutput[1], lexer.PeekCharAt(1).Rune)
	assert.Equal(t, helloOutput[2], lexer.PeekCharAt(2).Rune)

	for _, r := range helloOutput {
		assert.Equal(t, r, lexer.EatChar().Rune)
	}

	assert.Equal(t, rune(0), lexer.EatChar().Rune)
	assert.Equal(t, rune(0), lexer.EatChar().Rune)
	assert.Equal(t, rune(0), lexer.EatChar().Rune)

	assert.True(t, lexer.IsEof())
}

func TestReadCharsWithError(t *testing.T) {
	lexer := lexers.NewBaseLexer(helloInput, nil)
	lexer.MaxErrors = 2

	lexer.RegisterError("Error 1")
	lexer.RegisterError("Error 2")
	lexer.RegisterError("Error 3")

	assert.Equal(t, rune(0), lexer.EatChar().Rune)
	assert.True(t, lexer.IsEof())
}

func TestReadTokens(t *testing.T) {
	lexer := lexers.NewBaseLexer(helloInput, func(l *lexers.BaseLexer) *tokens.Token {
		c := l.EatChar()
		if c.Is(0) {
			return tokens.NewToken(tokens.EOF, "").WithRangeChars(c, l.PeekChar())
		}
		return tokens.NewToken(tokens.UNKNOWN, string(c.Rune)).WithRangeChars(c, l.PeekChar())
	})

	assert.Equal(t, string(helloOutput[0]), lexer.PeekToken().Literal)
	assert.Equal(t, string(helloOutput[1]), lexer.PeekTokenAt(1).Literal)
	assert.Equal(t, string(helloOutput[2]), lexer.PeekTokenAt(2).Literal)

	for i, r := range helloOutput {
		tk := lexer.EatToken()
		assert.Equal(t, 1, tk.FromLine)
		assert.Equal(t, i+1, tk.FromColumn)
		assert.Equal(t, string(r), tk.Literal)
		assert.Equal(t, tokens.TokenType(tokens.UNKNOWN), tk.Type)
	}

	assert.Equal(t, tokens.EOF, lexer.EatToken().Type)
	assert.Equal(t, tokens.EOF, lexer.EatToken().Type)
	assert.Equal(t, tokens.EOF, lexer.EatToken().Type)
}

func TestNext(t *testing.T) {
	lexer := lexers.NewBaseLexer(helloInput, func(l *lexers.BaseLexer) *tokens.Token {
		c := l.EatChar()
		if c.Is(0) {
			return tokens.NewToken(tokens.EOF, "").WithRangeChars(c, l.PeekChar())
		}
		return tokens.NewToken(tokens.UNKNOWN, string(c.Rune)).WithRangeChars(c, l.PeekChar())
	})

	i := 0
	for {
		tk, eof := lexer.Next()
		if eof {
			break
		}

		assert.Equal(t, string(helloOutput[i]), tk.Literal)
		i++
	}

	assert.True(t, lexer.IsEof())
}

func TestInterface(t *testing.T) {
	var lexer lexers.Lexer = lexers.NewBaseLexer([]byte{}, nil)
	assert.Empty(t, lexer.Errors())
}
