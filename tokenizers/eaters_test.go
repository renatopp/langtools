package tokenizers_test

import (
	"testing"

	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/tokenizers"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	input := []byte(`"hél'lo"`)
	expected := `hél'lo`

	l := langtools.NewLexer(input, func(l *langtools.Lexer) langtools.Token {
		return tokenizers.EatString(l, langtools.TString)
	})
	token, _ := l.Next()
	assert.Equal(t, expected, token.Literal)
}

func TestStringError(t *testing.T) {
	input := []byte(`"hél'lo`)

	l := langtools.NewLexer(input, func(l *langtools.Lexer) langtools.Token {
		return tokenizers.EatString(l, langtools.TString)
	})
	l.Next()
	assert.NotEmpty(t, l.Errors())
}

func TestRawString(t *testing.T) {
	input := []byte("`\"hello\nworld\"`")
	expected := "\"hello\nworld\""

	l := langtools.NewLexer(input, func(l *langtools.Lexer) langtools.Token {
		return tokenizers.EatRawString(l, langtools.TString)
	})
	token, _ := l.Next()
	assert.Equal(t, expected, token.Literal)
}

func TestRawStringError(t *testing.T) {
	input := []byte("`\"hello\nworld\"")

	l := langtools.NewLexer(input, func(l *langtools.Lexer) langtools.Token {
		return tokenizers.EatRawString(l, langtools.TString)
	})
	l.Next()
	assert.NotEmpty(t, l.Errors())
}
