package tokens_test

import (
	"testing"

	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

func TestCharAt(t *testing.T) {
	char := tokens.NewChar(1, 2, 3, 'a')
	line, column := char.At()
	assert.Equal(t, 1, line)
	assert.Equal(t, 2, column)
}

func TestCharIs(t *testing.T) {
	char := tokens.NewChar(1, 2, 3, 'a')
	assert.True(t, char.Is('a'))
}

func TestCharIsOneOf(t *testing.T) {
	char := tokens.NewChar(1, 2, 3, 'a')
	assert.True(t, char.IsOneOf('a', 'b', 'c'))
	assert.True(t, char.IsOneOf('c', 'b', 'a'))
}
