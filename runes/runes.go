package runes

import "slices"

// Check if a rune is a digit (0-9)
func IsDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// Check if a rune is a letter (a-z or A-Z)
func IsLetter(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
}

// Check if a rune is a whitespace (space, tab, newline or carriage return)
func IsWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
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
