package parsers

import (
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/tokens"
)

type PrattCurTokenFn func() tokens.Token
type PrattIsEndOfExprFn func(tokens.Token) bool
type PrattPrecedenceFn func(tokens.Token) int
type PrattPrefixFn func() asts.Node
type PrattInfixFn func(left asts.Node) asts.Node
type PrattPostfixFn func() asts.Node

type PrattParser struct {
	CurTokenFn   PrattCurTokenFn
	IsEndOfExpr  PrattIsEndOfExprFn
	PrecedenceFn PrattPrecedenceFn
	prefixFns    map[tokens.TokenType]PrattPrefixFn
	infixFns     map[tokens.TokenType]PrattInfixFn
	postfixFns   map[tokens.TokenType]PrattPostfixFn
}

type PrattParserOption func(*PrattParser)

func NewPrattParser(options ...PrattParserOption) *PrattParser {
	parser := &PrattParser{
		CurTokenFn:   func() tokens.Token { return tokens.Token{} },
		IsEndOfExpr:  func(tokens.Token) bool { return false },
		PrecedenceFn: func(tokens.Token) int { return 0 },
		prefixFns:    make(map[tokens.TokenType]PrattPrefixFn),
		infixFns:     make(map[tokens.TokenType]PrattInfixFn),
		postfixFns:   make(map[tokens.TokenType]PrattPostfixFn),
	}

	for _, option := range options {
		option(parser)
	}

	return parser
}

func (p *PrattParser) RegisterPrefixFn(tokenType tokens.TokenType, fn PrattPrefixFn) {
	p.prefixFns[tokenType] = fn
}

func (p *PrattParser) RegisterInfixFn(tokenType tokens.TokenType, fn PrattInfixFn) {
	p.infixFns[tokenType] = fn
}

func (p *PrattParser) RegisterPostfixFn(tokenType tokens.TokenType, fn PrattPostfixFn) {
	p.postfixFns[tokenType] = fn
}

func (p *PrattParser) ParseExpression(precedence int) asts.Node {
	postfix := p.postfixFns[p.CurTokenFn().Type]
	if postfix != nil {
		return postfix()
	}

	prefix := p.prefixFns[p.CurTokenFn().Type]
	if prefix == nil {
		return nil
	}

	left := prefix()

	c := p.CurTokenFn()
	for !p.IsEndOfExpr(c) && precedence < p.PrecedenceFn(c) {
		infix := p.infixFns[c.Type]
		if infix == nil {
			break
		}

		left = infix(left)
		c = p.CurTokenFn()
	}

	return left
}
