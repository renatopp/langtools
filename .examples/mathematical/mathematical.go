package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/parsers"
	"github.com/renatopp/langtools/runes"
	"github.com/renatopp/langtools/tokens"
)

type MathematicalLexer struct {
	*lexers.BaseLexer
}

func NewMathematicalLexer(input []byte) *MathematicalLexer {
	m := &MathematicalLexer{lexers.NewBaseLexer(input, nil)}
	m.TokenizerFn = func(bl *lexers.BaseLexer) tokens.Token {
		return m.tokenizer()
	}
	return m
}

func (l *MathematicalLexer) tokenizer() tokens.Token {
	for {
		if l.HasTooManyErrors() {
			break
		}

		c := l.PeekChar()
		switch {
		case runes.IsNumber(c.Rune):
			return l.EatNumber().As("number")

		case slices.Contains([]rune{'+', '-', '*', '/'}, c.Rune):
			l.EatChar()
			return tokens.NewTokenAtChar("operator", string(c.Rune), c)

		case runes.IsSpace(c.Rune):
			l.EatSpaces()

		case c.IsOneOf('\n', ';'):
			l.EatChar()
			return tokens.NewTokenAtChar("eoe", string(c.Rune), c)

		case runes.IsEof(c.Rune):
			return tokens.NewEofToken()

		default:
			l.EatChar()
			l.RegisterError("unexpected character")
		}
	}

	return tokens.NewEofToken()
}

type MathematicalParser struct {
	*parsers.PrattParser
}

func NewMathematicalParser(lexer *MathematicalLexer) *MathematicalParser {
	p := &MathematicalParser{}
	p.PrattParser = parsers.NewPrattParser(lexer)
	p.IsEndOfExpr = func(t tokens.Token) bool { return t.IsOneOfType(tokens.EOF, "eoe") }
	p.PrecedenceFn = func(t tokens.Token) int { return p.precedence(t) }
	p.RegisterPrefixFn("number", p.prefixNumber)
	return p
}

func (p *MathematicalParser) precedence(t tokens.Token) int {
	switch t.Literal {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}

func (p *MathematicalParser) prefixNumber() asts.Node {
	t := p.Lexer.EatToken()
	return &asts.BaseNode{Token: t}
}

func main() {
	input := []byte("1 + 2 * 3 / 4 - 5")

	// Test tokenizer
	{
		lexer := NewMathematicalLexer(input)
		tokens := lexer.All()
		for _, t := range tokens {
			println(t.DebugString())
		}
	}

	// Test parser
	{
		lexer := NewMathematicalLexer(input)
		parser := NewMathematicalParser(lexer)
		node := parser.ParseExpression(-1)

		println(node)
		node.Traverse(0, func(i int, n asts.Node) {
			fmt.Printf("%s%s\n", strings.Repeat("  ", i), n.String())
		})
	}
}
