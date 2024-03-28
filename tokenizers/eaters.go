package tokenizers

import (
	"strconv"

	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/runes"
)

func EatString(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""
	escaping := false
	first := l.EatChar()
	for {
		c := l.PeekChar()

		if c.Is('\n') {
			// Ignoring newlines
			l.RegisterError(langtools.ErrUnexpectedNewline)
			l.EatChar()
			continue

		} else if l.IsEof() {
			// Stopping at EOF
			l.RegisterError(langtools.ErrUnexpectedEndOfFile)
			break

		} else if !escaping && c.Is(first.Rune) {
			// End of string
			break

		} else if !escaping && c.Is('\\') {
			// Starting Escaping
			escaping = true
			l.EatChar()
			continue

		} else if escaping && !c.Is(first.Rune) {
			// Ending Escaping
			escaping = false
			r, err := strconv.Unquote(`"\` + string(c.Rune) + `"`)
			if err != nil {
				l.RegisterError(err.Error())
				l.EatChar()
				continue
			}
			c.Rune = rune(r[0])
		}

		result += string(c.Rune)
		l.EatChar()
	}

	l.EatChar()
	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatRawString(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""
	escaping := false
	first := l.EatChar()
	for {
		c := l.PeekChar()

		if l.IsEof() {
			// Stopping at EOF
			l.RegisterError(langtools.ErrUnexpectedEndOfFile)
			break

		} else if !escaping && c.Is(first.Rune) {
			// End of string
			break

		} else if !escaping && c.Is('\\') {
			// Starting Escaping
			escaping = true
			l.EatChar()
			continue

		} else if escaping && !c.Is(first.Rune) {
			// Ending Escaping
			escaping = false
			r, err := strconv.Unquote(`"\` + string(c.Rune) + `"`)
			if err != nil {
				l.RegisterError(err.Error())
				l.EatChar()
				continue
			}
			c.Rune = rune(r[0])
		}

		result += string(c.Rune)
		l.EatChar()
	}

	l.EatChar()
	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatNumber(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""
	dot := false
	exp := false

	first := l.PeekChar()
	for {
		c := l.PeekChar()

		if c.Is('.') {
			if dot || exp {
				l.RegisterError(langtools.ErrUnexpectedDot)
				l.EatChar()
				continue
			}

			dot = true
			result += string(c.Rune)

		} else if c.Is('e') || c.Is('E') {
			if exp {
				l.RegisterError(langtools.ErrUnexpectedE)
				l.EatChar()
				continue
			}

			exp = true
			result += string(c.Rune)

			next := l.PeekCharAt(1)
			if next.Is('+') || next.Is('-') {
				l.EatChar()
				result += string(next.Rune)
			}

		} else if runes.IsDigit(c.Rune) {
			result += string(c.Rune)

		} else {
			break
		}

		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatInteger(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if !runes.IsDigit(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatHexadecimal(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.EatChar()
	next := l.PeekChar()
	if first.Is('0') && (next.Is('x') || next.Is('X')) {
		l.EatChar()
	} else {
		result += string(first.Rune)
	}

	for {
		c := l.PeekChar()
		if !runes.IsHexadecimal(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatOctal(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.EatChar()
	if !first.Is('0') {
		result += string(first.Rune)
	}

	for {
		c := l.PeekChar()
		if !runes.IsOctal(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatBinary(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.EatChar()
	next := l.PeekChar()
	if first.Is('0') && (next.Is('b') || next.Is('B')) {
		l.EatChar()
	} else {
		result += string(first.Rune)
	}

	for {
		c := l.PeekChar()
		if !runes.IsBinary(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatWhitespaces(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if !runes.IsWhitespace(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatNewlines(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if !c.Is('\n') {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatIdentifier(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if !runes.IsAlphaNumeric(c.Rune) && !c.Is('_') {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatWord(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if !runes.IsAlphaNumeric(c.Rune) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}

func EatUntilEndOfLine(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if c.Is('\n') || c.Is(0) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return langtools.NewToken(tp, result, first.Line, first.Column)
}
