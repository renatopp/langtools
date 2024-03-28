package tokenizers

import (
	"strconv"

	"github.com/renatopp/langtools"
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
		println("processing", string(c.Rune))

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
	return langtools.NewToken(tp, "", 0, 0)
}

func EatInteger(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatHexadecimal(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatOctal(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatBinary(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatWhitespaces(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatNewlines(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatIdentifier(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatBacklash(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatWorld(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}

func EatComment(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
	return langtools.NewToken(tp, "", 0, 0)
}
