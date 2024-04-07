package utils

import "strings"

// PadLeft pads the string with spaces on the left side to the specified length.
func PadLeft(str string, length int) string {
	return PadLeftWith(str, length, " ")
}

// PadLeftWith pads the string with the specified character on the left side to
// the specified length.
func PadLeftWith(str string, length int, char string) string {
	if len(str) >= length {
		return str
	}

	return strings.Repeat(char, length-len(str)) + str
}

// PadRight pads the string with spaces on the right side to the specified
// length.
func PadRight(str string, length int) string {
	return PadRightWith(str, length, " ")
}

// PadRightWith pads the string with the specified character on the right side
// to the specified length.
func PadRightWith(str string, length int, char string) string {
	for len(str) < length {
		str = str + char
	}
	return str
}

// PadCenter pads the string with spaces on both sides to the specified length.
func PadCenter(str string, length int) string {
	return PadCenterWith(str, length, " ")
}

// PadCenterWith pads the string with the specified character on both sides to
// the specified length.
func PadCenterWith(str string, length int, char string) string {
	if len(str) >= length {
		return str
	}

	left := (length - len(str)) / 2
	right := length - len(str) - left

	return strings.Repeat(char, left) + str + strings.Repeat(char, right)
}

// GetSourceLines returns the lines of the source code from the specified
// range. Notice that lines starts at 1.
func GetSourceLines(source []byte, from, to int) []string {
	lines := strings.Split(string(source), "\n")
	return lines[from-1 : to]
}

// HighlightChars returns a string with the specified characters highlighted.
// See HighlightCharsWith for more information.
func HighlightChars(from, to int) string {
	return HighlightCharsWith(from, to, "^")
}

// HighlightCharsWith returns a string with the specified characters
// highlighted with the specified character. Notice that the characters starts
// at 1.
//
// This is a simple helper function to help highlighting the characters or
// tokens in the source code that are causing the error. The output of this
// function is like:
//
//     HighlightCharsWith(2, 5, "^")
//     // Output: " ^^^"
func HighlightCharsWith(from, to int, char string) string {
	return strings.Repeat(" ", from-1) + strings.Repeat(char, to-from)
}
