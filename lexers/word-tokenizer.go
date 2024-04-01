package lexers

import (
	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/tokens"
)

// Simple word tokenizer. It ignore whitespaces, group punctuation, and return
// words and numbers.
func WordTokenizer(l *BaseLexer) tokens.Token {
	for {
		if l.HasTooManyErrors() {
			break
		}

		c := l.PeekChar()
		switch {
		case c.Is(0):
			return tokens.NewEofTokenAtChar(c)

		case runes.IsWhitespace(c.Rune):
			l.EatWhitespaces()

		case runes.IsLetter(c.Rune):
			return l.EatWord().As("word")

		case runes.IsDigit(c.Rune):
			return l.EatNumber().As("number")

		default:
			return eatPunctuation(l).As("punctuation")
		}
	}

	return tokens.NewEofToken()
}

func eatPunctuation(l *BaseLexer) tokens.Token {
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

	return tokens.NewTokenAtChar(tokens.UNKNOWN, result, first)
}
