package utils_test

import (
	"testing"

	"github.com/renatopp/langtools/utils"
	"github.com/stretchr/testify/assert"
)

func TestPadLeft(t *testing.T) {
	str := "test"
	result := utils.PadLeft(str, 10)
	expected := "      test"
	assert.Equal(t, expected, result)
}

func TestPadLeftWith(t *testing.T) {
	str := ".another."
	result := utils.PadLeftWith(str, 12, ".")
	expected := "....another."
	assert.Equal(t, expected, result)
}

func TestPadRight(t *testing.T) {
	str := "test"
	result := utils.PadRight(str, 10)
	expected := "test      "
	assert.Equal(t, expected, result)
}

func TestPadRightWith(t *testing.T) {
	str := ".another."
	result := utils.PadRightWith(str, 12, ".")
	expected := ".another...."
	assert.Equal(t, expected, result)
}

func TestPadCenter(t *testing.T) {
	str := "test"
	result := utils.PadCenter(str, 10)
	expected := "   test   "
	assert.Equal(t, expected, result)
}

func TestPadCenterWith(t *testing.T) {
	str := ".another."
	result := utils.PadCenterWith(str, 12, ".")
	expected := "..another..."
	assert.Equal(t, expected, result)
}

func TestGetSourceLines(t *testing.T) {
	source := []byte("line1\nline2\nline3\nline4\nline5\nline6\nline7")
	result := utils.GetSourceLines(source, 2, 4)
	expected := []string{"line2", "line3", "line4"}
	assert.Equal(t, expected, result)
}

func TestHighlightChars(t *testing.T) {
	expected := "     ^^^^"
	result := utils.HighlightChars(6, 10)
	assert.Equal(t, expected, result)
}

func TestHighlightCharsWith(t *testing.T) {
	expected := "     ---------"
	result := utils.HighlightCharsWith(6, 15, "-")
	assert.Equal(t, expected, result)
}
