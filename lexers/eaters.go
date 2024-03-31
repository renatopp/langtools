package lexers

import (
	"strconv"

	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/token"
)

func EatString(l *GenericLexer, tp token.TokenType) token.Token {
	result := ""
	escaping := false
	first := l.EatChar()
	for {
		c := l.PeekChar()

		if c.Is('\n') {
			// Ignoring newlines
			l.RegisterError(ErrUnexpectedNewline)
			l.EatChar()
			continue

		} else if l.IsEof() {
			// Stopping at EOF
			l.RegisterError(ErrUnexpectedEndOfFile)
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
	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatRawString(l *GenericLexer, tp token.TokenType) token.Token {
	result := ""
	escaping := false
	first := l.EatChar()
	for {
		c := l.PeekChar()

		if l.IsEof() {
			// Stopping at EOF
			l.RegisterError(ErrUnexpectedEndOfFile)
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
	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatNumber(l *GenericLexer, tp token.TokenType) token.Token {
	result := ""
	dot := false
	exp := false

	first := l.PeekChar()
	for {
		c := l.PeekChar()

		if c.Is('.') {
			if dot || exp {
				l.RegisterError(ErrUnexpectedDot)
				l.EatChar()
				continue
			}

			dot = true
			result += string(c.Rune)

		} else if c.Is('e') || c.Is('E') {
			if exp {
				l.RegisterError(ErrUnexpectedE)
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatInteger(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatHexadecimal(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatOctal(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatBinary(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatWhitespaces(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatNewlines(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatIdentifier(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatWord(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}

func EatUntilEndOfLine(l *GenericLexer, tp token.TokenType) token.Token {
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

	return token.NewToken(tp, result, first.Line, first.Column)
}
