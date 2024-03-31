package lexers

import (
	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/token"
)

const (
	TWord        = "word"
	TNumber      = "number"
	TPunctuation = "punctuation"
)

// Simple word tokenizer. It ignore whitespaces, group punctuation, and return
// words and numbers.
func Word(l *Lexer) token.Token {
	for {
		if l.HasTooManyErrors() {
			break
		}

		c := l.PeekChar()
		switch {
		case c.Is(0):
			return token.NewToken(token.TEof, "", c.Line, c.Column)

		case runes.IsWhitespace(c.Rune):
			EatWhitespaces(l, token.TString)

		case runes.IsLetter(c.Rune):
			return EatWord(l, TWord)

		case runes.IsDigit(c.Rune):
			return EatNumber(l, TNumber)

		default:
			return eatPunctuation(l, TPunctuation)

		}
	}

	return token.Token{}
}

func eatPunctuation(l *Lexer, tp token.TokenType) token.Token {
	result := ""

	first := l.PeekChar()
	for {
		c := l.PeekChar()
		if runes.IsAlphaNumeric(c.Rune) || runes.IsWhitespace(c.Rune) || c.Is(0) {
			break
		}

		result += string(c.Rune)
		l.EatChar()
	}

	return token.NewToken(tp, result, first.Line, first.Column)
}
