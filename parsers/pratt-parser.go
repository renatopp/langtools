package parsers

import (
	"github.com/renatopp/langtools/asts"
	"github.com/renatopp/langtools/lexers"
	"github.com/renatopp/langtools/tokens"
)

type PrattCurTokenFn func() tokens.Token
type PrattIsEndOfExprFn func(tokens.Token) bool
type PrattPrecedenceFn func(tokens.Token) int
type PrattPrefixFn func() asts.Node
type PrattInfixFn func(left asts.Node) asts.Node
type PrattPostfixFn func() asts.Node

type PrattParser struct {
	*BaseParser
	IsEndOfExpr  PrattIsEndOfExprFn
	PrecedenceFn PrattPrecedenceFn
	prefixFns    map[tokens.TokenType]PrattPrefixFn
	infixFns     map[tokens.TokenType]PrattInfixFn
	postfixFns   map[tokens.TokenType]PrattPostfixFn
}

type PrattParserOption func(*PrattParser)

func NewPrattParser(lexer lexers.Lexer, options ...PrattParserOption) *PrattParser {
	parser := &PrattParser{
		BaseParser:   NewBaseParser(lexer),
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
	postfix := p.postfixFns[p.Lexer.PeekToken().Type]
	if postfix != nil {
		return postfix()
	}

	prefix := p.prefixFns[p.Lexer.PeekToken().Type]
	if prefix == nil {
		return nil
	}

	left := prefix()

	c := p.Lexer.PeekToken()
	for !p.IsEndOfExpr(c) && precedence < p.PrecedenceFn(c) {
		infix := p.infixFns[c.Type]
		if infix == nil {
			break
		}

		left = infix(left)
		c = p.Lexer.PeekToken()
	}

	return left
}
