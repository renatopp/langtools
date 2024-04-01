package runes

import "slices"

// Check if a rune is a digit (0-9)
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Alias for IsDigit
var IsNumber = IsDigit

// Alias for IsDigit
var IsNumeric = IsDigit

// Check if a rune is a letter (a-z or A-Z)
func IsLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

// Alias for IsLetter
var IsAlpha = IsLetter

// Check if a rune is a whitespace (space, tab, newline or carriage return)
func IsWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// Check if a rune is a space (space, tab or carriage return)
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

// Check if a rune is the end of file (rune 0)
func IsEof(r rune) bool {
	return r == 0
}

// Check if a rune is a letter or a digit
func IsAlphaNumeric(r rune) bool {
	return IsLetter(r) || IsDigit(r)
}

// Check if a rune is any of the runes in the given slice
func IsAnyOf(r rune, runes []rune) bool {
	return slices.Contains(runes, r)
}

// Check if a rune is a hexadecimal digit (0-9, a-f or A-F)
func IsHexadecimal(r rune) bool {
	return IsDigit(r) || r >= 'a' && r <= 'f' || r >= 'A' && r <= 'F'
}

// Check if a rune is an octal digit (0-7)
func IsOctal(r rune) bool {
	return r >= '0' && r <= '7'
}

// Check if a rune is a binary digit (0 or 1)
func IsBinary(r rune) bool {
	return r == '0' || r == '1'
}
