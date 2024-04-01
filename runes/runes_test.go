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
	assert.True(t, runes.IsWhiteSpace(' '))
	assert.True(t, runes.IsWhiteSpace('\t'))
	assert.True(t, runes.IsWhiteSpace('\n'))
	assert.True(t, runes.IsWhiteSpace('\r'))
	assert.False(t, runes.IsWhiteSpace('a'))
	assert.False(t, runes.IsWhiteSpace('A'))
	assert.False(t, runes.IsWhiteSpace('0'))
	assert.False(t, runes.IsWhiteSpace('9'))
	assert.False(t, runes.IsWhiteSpace(';'))
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

func TestIsHexadecimal(t *testing.T) {
	assert.True(t, runes.IsHexadecimal('0'))
	assert.True(t, runes.IsHexadecimal('1'))
	assert.True(t, runes.IsHexadecimal('2'))
	assert.True(t, runes.IsHexadecimal('3'))
	assert.True(t, runes.IsHexadecimal('4'))
	assert.True(t, runes.IsHexadecimal('5'))
	assert.True(t, runes.IsHexadecimal('6'))
	assert.True(t, runes.IsHexadecimal('7'))
	assert.True(t, runes.IsHexadecimal('8'))
	assert.True(t, runes.IsHexadecimal('9'))
	assert.True(t, runes.IsHexadecimal('a'))
	assert.True(t, runes.IsHexadecimal('b'))
	assert.True(t, runes.IsHexadecimal('c'))
	assert.True(t, runes.IsHexadecimal('d'))
	assert.True(t, runes.IsHexadecimal('e'))
	assert.True(t, runes.IsHexadecimal('f'))
	assert.True(t, runes.IsHexadecimal('A'))
	assert.True(t, runes.IsHexadecimal('B'))
	assert.True(t, runes.IsHexadecimal('C'))
	assert.True(t, runes.IsHexadecimal('D'))
	assert.True(t, runes.IsHexadecimal('E'))
	assert.True(t, runes.IsHexadecimal('F'))
	assert.False(t, runes.IsHexadecimal(' '))
	assert.False(t, runes.IsHexadecimal(';'))
}

func TestIsOctal(t *testing.T) {
	assert.True(t, runes.IsOctal('0'))
	assert.True(t, runes.IsOctal('1'))
	assert.True(t, runes.IsOctal('2'))
	assert.True(t, runes.IsOctal('3'))
	assert.True(t, runes.IsOctal('4'))
	assert.True(t, runes.IsOctal('5'))
	assert.True(t, runes.IsOctal('6'))
	assert.True(t, runes.IsOctal('7'))
	assert.False(t, runes.IsOctal('8'))
	assert.False(t, runes.IsOctal('9'))
	assert.False(t, runes.IsOctal('a'))
	assert.False(t, runes.IsOctal('A'))
	assert.False(t, runes.IsOctal(' '))
	assert.False(t, runes.IsOctal(';'))
}

func TestIsBinary(t *testing.T) {
	assert.True(t, runes.IsBinary('0'))
	assert.True(t, runes.IsBinary('1'))
	assert.False(t, runes.IsBinary('2'))
	assert.False(t, runes.IsBinary('3'))
	assert.False(t, runes.IsBinary('4'))
	assert.False(t, runes.IsBinary('5'))
	assert.False(t, runes.IsBinary('6'))
	assert.False(t, runes.IsBinary('7'))
	assert.False(t, runes.IsBinary('8'))
	assert.False(t, runes.IsBinary('9'))
	assert.False(t, runes.IsBinary('a'))
	assert.False(t, runes.IsBinary('A'))
	assert.False(t, runes.IsBinary(' '))
	assert.False(t, runes.IsBinary(';'))
}
