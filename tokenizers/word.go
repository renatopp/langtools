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

func Word(l *langtools.Lexer) langtools.Token {
	var token *langtools.Token

	for {
		if l.HasTooManyErrors() {
			break
		}

		c := l.PeekChar()
		switch {
		case runes.IsWhitespace(c.Rune):
			l.EatChar()
		case runes.IsLetter(c.Rune):
			s := eatWord(l)
			return langtools.NewToken(TWord, s, c.Line, c.Column)
		case runes.IsDigit(c.Rune):
			s := eatNumber(l)
			return langtools.NewToken(TNumber, s, c.Line, c.Column)
		default:
			s := eatPunctuation(l)
			return langtools.NewToken(TPunctuation, s, c.Line, c.Column)
		}
	}

	return *token
}

func eatWord(l *langtools.Lexer) string {
	c := l.EatChar()
	word := string(c.Rune)
	for {
		c = l.PeekChar()
		if !runes.IsLetter(c.Rune) {
			break
		}

		word += string(c.Rune)
		l.EatChar()
	}

	return word
}

func eatNumber(l *langtools.Lexer) string {
	c := l.EatChar()
	number := string(c.Rune)
	for {
		c = l.PeekChar()
		if !runes.IsDigit(c.Rune) {
			break
		}

		number += string(c.Rune)
		l.EatChar()
	}

	return number
}

func eatPunctuation(l *langtools.Lexer) string {
	c := l.EatChar()
	punc := string(c.Rune)
	for {
		c = l.PeekChar()
		if runes.IsAlphaNumeric(c.Rune) || runes.IsWhitespace(c.Rune) || runes.IsEof(c.Rune) {
			break
		}

		punc += string(c.Rune)
		l.EatChar()
	}
	return punc
}
