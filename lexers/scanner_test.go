package lexers_test

import (
	"testing"

	"github.com/renatopp/langtools/lexers"
	"github.com/stretchr/testify/assert"
)

func TestScanner(t *testing.T) {
	input := []byte("Hello, World!")
	expected := []rune("Hello, World!")
	scanner := lexers.NewScanner(input)

	for i, r := range expected {
		char, _ := scanner.Next()
		if char.Rune == 0 {
			break
		}

		assert.Equal(t, char.Column, i+1)
		assert.Equal(t, char.Rune, r)
	}

	char, _ := scanner.Next()
	assert.Equal(t, char.Rune, rune(0))

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, rune(0))

	assert.True(t, scanner.IsEof())
}

func TestScannerAll(t *testing.T) {
	input := []byte("Hello, World!")
	expected := []rune("Hello, World!")
	scanner := lexers.NewScanner(input)

	chars, _ := scanner.All()
	for i, r := range expected {
		assert.Equal(t, chars[i].Column, i+1)
		assert.Equal(t, chars[i].Rune, r)
	}

	assert.Equal(t, len(chars), len(expected))

	char, _ := scanner.Next()
	assert.Equal(t, char.Rune, rune(0))
}

func TestScannerInvalidChar(t *testing.T) {
	input := []byte("Hello, \xFF World!")
	scanner := lexers.NewScanner(input)

	res, err := scanner.All()
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestReset(t *testing.T) {
	input := []byte("Hello, World!")
	scanner := lexers.NewScanner(input)

	char, _ := scanner.Next()
	assert.Equal(t, char.Rune, 'H')

	scanner.Reset()

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'H')
}

func TestScannerNewLine(t *testing.T) {
	input := []byte("Hello,\nWorld!")
	scanner := lexers.NewScanner(input)

	char, _ := scanner.Next()
	assert.Equal(t, char.Rune, 'H')
	assert.Equal(t, char.Column, 1)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'e')
	assert.Equal(t, char.Column, 2)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'l')
	assert.Equal(t, char.Column, 3)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'l')
	assert.Equal(t, char.Column, 4)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'o')
	assert.Equal(t, char.Column, 5)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, ',')
	assert.Equal(t, char.Column, 6)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, '\n')
	assert.Equal(t, char.Column, 7)
	assert.Equal(t, char.Line, 1)

	char, _ = scanner.Next()
	assert.Equal(t, char.Rune, 'W')
	assert.Equal(t, char.Column, 1)
	assert.Equal(t, char.Line, 2)
}
