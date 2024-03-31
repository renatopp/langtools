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
}
