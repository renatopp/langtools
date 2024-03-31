package lexers_test

import (
	"testing"

	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/token"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	input := []byte(`"hél'lo"`)
	expected := `hél'lo`

	l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
		return lexers.EatString(l, token.TString)
	})
	token, _ := l.Next()
	assert.Equal(t, expected, token.Literal)
}

func TestStringError(t *testing.T) {
	input := []byte(`"hél'lo`)

	l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
		return lexers.EatString(l, token.TString)
	})
	l.Next()
	assert.NotEmpty(t, l.Errors())
}

func TestRawString(t *testing.T) {
	input := []byte("`\"hello\nworld\"`")
	expected := "\"hello\nworld\""

	l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
		return lexers.EatRawString(l, token.TString)
	})
	token, _ := l.Next()
	assert.Equal(t, expected, token.Literal)
}

func TestRawStringError(t *testing.T) {
	input := []byte("`\"hello\nworld\"")

	l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
		return lexers.EatRawString(l, token.TString)
	})
	l.Next()
	assert.NotEmpty(t, l.Errors())
}

func TestNumber(t *testing.T) {
	inputs := [][]byte{
		[]byte(`123!`),
		[]byte(`123.456]`),
		[]byte(`123e456,`),
		[]byte(`123.456e789`),
		[]byte(`123e+16`),
		[]byte(`123.2e-16`),
		[]byte(`.123`),
	}
	expected := []string{
		`123`,
		`123.456`,
		`123e456`,
		`123.456e789`,
		`123e+16`,
		`123.2e-16`,
		`.123`,
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatNumber(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestNumberError(t *testing.T) {
	inputs := [][]byte{
		[]byte(`123.2.3`),
		[]byte(`123e12.3`),
		[]byte(`123e12e2`),
	}

	for _, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatNumber(l, token.TString)
		})
		l.Next()
		assert.NotEmpty(t, l.Errors())
	}
}

func TestInteger(t *testing.T) {
	inputs := [][]byte{
		[]byte(`123`),
		[]byte(`123.2`),
		[]byte(`123e3`),
		[]byte(`123|123`),
	}
	expected := []string{
		`123`,
		`123`,
		`123`,
		`123`,
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatInteger(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestHexadecimal(t *testing.T) {
	inputs := [][]byte{
		[]byte(`0x123`),
		[]byte(`0XAAFFDD`),
		[]byte(`0x1.2`),
		[]byte(`0x1G`),
		[]byte(`AAFFEE`),
	}
	expected := []string{
		`123`,
		`AAFFDD`,
		`1`,
		`1`,
		`AAFFEE`,
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatHexadecimal(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestOctal(t *testing.T) {
	inputs := [][]byte{
		[]byte(`0770`),
		[]byte(`0012`),
		[]byte(`0120`),
		[]byte(`123`),
		[]byte(`189`),
	}
	expected := []string{
		`770`,
		`012`,
		`120`,
		`123`,
		`1`,
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatOctal(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestBinary(t *testing.T) {
	inputs := [][]byte{
		[]byte(`0b11000000110`),
		[]byte(`0B110`),
		[]byte(`100110`),
		[]byte(`1002`),
	}
	expected := []string{
		`11000000110`,
		`110`,
		`100110`,
		`100`,
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatBinary(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestWhitespaces(t *testing.T) {
	inputs := [][]byte{
		[]byte("\n \n"),
		[]byte("\n\t\t"),
	}
	expected := []string{
		"\n \n",
		"\n\t\t",
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatWhitespaces(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestNewlines(t *testing.T) {
	inputs := [][]byte{
		[]byte("\n\n"),
		[]byte("\n\t\t"),
	}
	expected := []string{
		"\n\n",
		"\n",
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatNewlines(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestIdentifier(t *testing.T) {
	inputs := [][]byte{
		[]byte("abc123"),
		[]byte("sample_variable"),
		[]byte("9var"),
		[]byte("var"),
	}
	expected := []string{
		"abc123",
		"sample_variable",
		"9var",
		"var",
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatIdentifier(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestWord(t *testing.T) {
	inputs := [][]byte{
		[]byte("renato"),
		[]byte("Renato"),
		[]byte("Hello World"),
		[]byte("Hello, World"),
		[]byte("B33 "),
	}
	expected := []string{
		"renato",
		"Renato",
		"Hello",
		"Hello",
		"B33",
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatWord(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}

func TestUntilEndOfLine(t *testing.T) {
	inputs := [][]byte{
		[]byte("renato\nteste"),
		[]byte("Renato abc \n123"),
		[]byte("Hello World"),
		[]byte("Hello,\n World"),
		[]byte("a\\na"),
		[]byte("a\\\na"),
	}
	expected := []string{
		"renato",
		"Renato abc ",
		"Hello World",
		"Hello,",
		"a\\na",
		"a\\",
	}

	for i, input := range inputs {
		l := lexers.NewGenericLexer(input, func(l *lexers.GenericLexer) token.Token {
			return lexers.EatUntilEndOfLine(l, token.TNumber)
		})
		token, _ := l.Next()
		assert.Equal(t, expected[i], token.Literal)
	}
}