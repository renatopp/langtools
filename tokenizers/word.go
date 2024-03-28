package tokenizers

import (
	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/runes"
)

const (
	TWord        = "word"
	TNumber      = "number"
	TPunctuation = "punctuation"
)

// Simple word tokenizer. It ignore whitespaces, group punctuation, and return
// words and numbers.
func Word(l *langtools.Lexer) langtools.Token {
	var token *langtools.Token

	for {
		if l.HasTooManyErrors() {
			break
		}

		c := l.PeekChar()
		switch {
		case runes.IsWhitespace(c.Rune):
			EatWhitespaces(l, langtools.TString)

		case runes.IsLetter(c.Rune):
			return EatWord(l, TWord)

		case runes.IsDigit(c.Rune):
			return EatNumber(l, TNumber)

		default:
			return eatPunctuation(l, TPunctuation)

		}
	}

	return *token
}

func eatPunctuation(l *langtools.Lexer, tp langtools.TokenType) langtools.Token {
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

	return langtools.NewToken(tp, result, first.Line, first.Column)
}
