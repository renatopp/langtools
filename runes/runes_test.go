package runes_test

import (
	"testing"

	"github.com/renatopp/langtools/runes"
	"github.com/stretchr/testify/assert"
)

func TestIsDigit(t *testing.T) {
	assert.True(t, runes.IsDigit('0'))
	assert.True(t, runes.IsDigit('1'))
	assert.True(t, runes.IsDigit('2'))
	assert.True(t, runes.IsDigit('3'))
	assert.True(t, runes.IsDigit('4'))
	assert.True(t, runes.IsDigit('5'))
	assert.True(t, runes.IsDigit('6'))
	assert.True(t, runes.IsDigit('7'))
	assert.True(t, runes.IsDigit('8'))
	assert.True(t, runes.IsDigit('9'))
	assert.False(t, runes.IsDigit('a'))
	assert.False(t, runes.IsDigit('A'))
	assert.False(t, runes.IsDigit(' '))
	assert.False(t, runes.IsDigit(';'))
}

func TestIsLetter(t *testing.T) {
	assert.True(t, runes.IsLetter('a'))
	assert.True(t, runes.IsLetter('b'))
	assert.True(t, runes.IsLetter('c'))
	assert.True(t, runes.IsLetter('A'))
	assert.True(t, runes.IsLetter('B'))
	assert.True(t, runes.IsLetter('W'))
	assert.True(t, runes.IsLetter('X'))
	assert.False(t, runes.IsLetter('0'))
	assert.False(t, runes.IsLetter('9'))
	assert.False(t, runes.IsLetter(' '))
	assert.False(t, runes.IsLetter(';'))
}

func TestIsWhitespace(t *testing.T) {
	assert.True(t, runes.IsWhitespace(' '))
	assert.True(t, runes.IsWhitespace('\t'))
	assert.True(t, runes.IsWhitespace('\n'))
	assert.True(t, runes.IsWhitespace('\r'))
	assert.False(t, runes.IsWhitespace('a'))
	assert.False(t, runes.IsWhitespace('A'))
	assert.False(t, runes.IsWhitespace('0'))
	assert.False(t, runes.IsWhitespace('9'))
	assert.False(t, runes.IsWhitespace(';'))
}

func TestIsEof(t *testing.T) {
	assert.True(t, runes.IsEof(0))
	assert.False(t, runes.IsEof('a'))
	assert.False(t, runes.IsEof('A'))
	assert.False(t, runes.IsEof('0'))
	assert.False(t, runes.IsEof('9'))
	assert.False(t, runes.IsEof(';'))
}

func TestIsAlphaNumeric(t *testing.T) {
	assert.True(t, runes.IsAlphaNumeric('a'))
	assert.True(t, runes.IsAlphaNumeric('A'))
	assert.True(t, runes.IsAlphaNumeric('0'))
	assert.True(t, runes.IsAlphaNumeric('9'))
	assert.False(t, runes.IsAlphaNumeric(' '))
	assert.False(t, runes.IsAlphaNumeric(';'))
}

func TestIsAnyOf(t *testing.T) {
	assert.True(t, runes.IsAnyOf('a', []rune{'a', 'b', 'c'}))
	assert.True(t, runes.IsAnyOf('b', []rune{'a', 'b', 'c'}))
	assert.True(t, runes.IsAnyOf('c', []rune{'a', 'b', 'c'}))
	assert.False(t, runes.IsAnyOf('d', []rune{'a', 'b', 'c'}))
	assert.False(t, runes.IsAnyOf(' ', []rune{'a', 'b', 'c'}))
	assert.False(t, runes.IsAnyOf(';', []rune{'a', 'b', 'c'}))
}
