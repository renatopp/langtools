package tokens_test

import (
	"testing"

	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

func TestTokenCreation(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithRange(1, 2, 3, 4)

	assert.Equal(t, tokens.EOF, token.Type)
	assert.Equal(t, "", token.Literal)
	assert.Equal(t, 1, token.FromLine)
	assert.Equal(t, 2, token.FromColumn)
	assert.Equal(t, 3, token.ToLine)
	assert.Equal(t, 4, token.ToColumn)
}

func TestTokenCreationWithChars(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithRangeChars(
		tokens.NewChar(1, 2, 1, '1'),
		tokens.NewChar(3, 4, 1, '1'),
	)

	assert.Equal(t, tokens.EOF, token.Type)
	assert.Equal(t, "", token.Literal)
	assert.Equal(t, 1, token.FromLine)
	assert.Equal(t, 2, token.FromColumn)
	assert.Equal(t, 3, token.ToLine)
	assert.Equal(t, 4, token.ToColumn)
}

func TestTokenWithType(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithType(tokens.INVALID)

	assert.Equal(t, tokens.INVALID, token.Type)
	assert.Equal(t, "", token.Literal)
}

func TestTokenWithLiteral(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithLiteral("literal")

	assert.Equal(t, tokens.EOF, token.Type)
	assert.Equal(t, "literal", token.Literal)
}

func TestRanges(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithRange(1, 2, 3, 4)

	fromLine, toLine := token.RangeLines()
	assert.Equal(t, 1, fromLine)
	assert.Equal(t, 3, toLine)

	fromColumn, toColumn := token.RangeColumns()
	assert.Equal(t, 2, fromColumn)
	assert.Equal(t, 4, toColumn)

	fromLine, fromColumn, toLine, toColumn = token.Range()
	assert.Equal(t, 1, fromLine)
	assert.Equal(t, 3, toLine)
	assert.Equal(t, 2, fromColumn)
	assert.Equal(t, 4, toColumn)
}

func TestIsType(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "")

	assert.True(t, token.IsType(tokens.EOF))
	assert.False(t, token.IsType(tokens.INVALID))
}

func TestIsLiteral(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "")

	assert.True(t, token.IsLiteral(""))
	assert.False(t, token.IsLiteral("literal"))
}

func TestIsOneOfTypes(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "")

	assert.True(t, token.IsOneOfTypes(tokens.EOF, tokens.INVALID))
	assert.False(t, token.IsOneOfTypes(tokens.INVALID))
}

func TestIsOneOfLiterals(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "")

	assert.True(t, token.IsOneOfLiterals("", "literal"))
	assert.False(t, token.IsOneOfLiterals("literal"))
}

func TestFrom(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithRange(1, 2, 3, 4)

	fromLine, fromColumn := token.From()
	assert.Equal(t, 1, fromLine)
	assert.Equal(t, 2, fromColumn)
}

func TestTo(t *testing.T) {
	token := tokens.NewToken(tokens.EOF, "").WithRange(1, 2, 3, 4)

	toLine, toColumn := token.To()
	assert.Equal(t, 3, toLine)
	assert.Equal(t, 4, toColumn)
}
