package lexers_test

import (
	"reflect"
	"testing"

	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/tokens"
	"github.com/stretchr/testify/assert"
)

func testValid(t *testing.T, inputs [][]byte, expects []string, eater string) {
	for i, expected := range expects {
		l := lexers.NewBaseLexer(inputs[i], nil)
		token := reflect.ValueOf(l).MethodByName(eater).Call([]reflect.Value{})[0].Interface().(*tokens.Token)
		assert.Equal(t, expected, token.Literal)
	}
}

func testError(t *testing.T, inputs [][]byte, eater string) {
	for _, input := range inputs {
		l := lexers.NewBaseLexer(input, nil)
		reflect.ValueOf(l).MethodByName(eater).Call([]reflect.Value{})
		assert.NotEmpty(t, l.Errors())
	}
}

func TestString(t *testing.T) {
	input := [][]byte{
		[]byte(`"hello"`),
		[]byte(`'hél"lo'`),
		[]byte(`/hél'lo/`),
	}
	expected := []string{
		`hello`,
		`hél"lo`,
		`hél'lo`,
	}

	testValid(t, input, expected, "EatString")
}

func TestStringError(t *testing.T) {
	input := [][]byte{
		[]byte(`"hél'lo`),
		[]byte(`'hél"lo`),
		[]byte(`/hél'lo`),
	}

	testError(t, input, "EatString")
}

func TestRawString(t *testing.T) {
	input := [][]byte{
		[]byte(`"hello
"`),
	}
	expected := []string{
		"hello\n",
	}

	testValid(t, input, expected, "EatRawString")
}

func TestRawStringError(t *testing.T) {
	input := [][]byte{
		[]byte(`"hél'lo`),
		[]byte(`'hél"lo`),
		[]byte(`/hél'lo`),
	}

	testError(t, input, "EatRawString")
}

func TestNumber(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatNumber")
}

func TestNumberError(t *testing.T) {
	input := [][]byte{
		[]byte(`123.2.3`),
		[]byte(`123e12.3`),
		[]byte(`123e12e2`),
	}

	testError(t, input, "EatNumber")
}

func TestInteger(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatInteger")
}

func TestHexadecimal(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatHexadecimal")
}

func TestOctal(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatOctal")
}

func TestBinary(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatBinary")
}

func TestWhitespaces(t *testing.T) {
	input := [][]byte{
		[]byte("\n \n"),
		[]byte("\n\t\t"),
	}
	expected := []string{
		"\n \n",
		"\n\t\t",
	}

	testValid(t, input, expected, "EatWhitespaces")
}

func TestSpaces(t *testing.T) {
	input := [][]byte{
		[]byte(" \n"),
		[]byte("\t\t\r\n"),
	}
	expected := []string{
		" ",
		"\t\t\r",
	}

	testValid(t, input, expected, "EatSpaces")
}

func TestNewlines(t *testing.T) {
	input := [][]byte{
		[]byte("\n\n"),
		[]byte("\n\t\t"),
	}
	expected := []string{
		"\n\n",
		"\n",
	}

	testValid(t, input, expected, "EatNewlines")
}

func TestIdentifier(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatIdentifier")
}

func TestWord(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatWord")
}

func TestUntilEndOfLine(t *testing.T) {
	input := [][]byte{
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

	testValid(t, input, expected, "EatUntilEndOfLine")
}
