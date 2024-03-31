package parsers

import (
	"github.com/renatopp/langtools/ast"
	"github.com/renatopp/langtools/token"
)

type PrattPrecedenceFn func(token.Token) int
type PrattPrefixFn func() ast.Node
type PrattInfixFn func(left ast.Node) ast.Node
type PrattPostfixFn func(left ast.Node) ast.Node

type PrattParser struct {
	precedenceFn PrattPrecedenceFn
	prefixFns    map[token.TokenType]PrattPrefixFn
	infixFns     map[token.TokenType]PrattInfixFn
	postfixFns   map[token.TokenType]PrattPostfixFn
}

func NewPrattParser(precedenceFn PrattPrecedenceFn) *PrattParser {
	return &PrattParser{
		precedenceFn: precedenceFn,
		prefixFns:    make(map[token.TokenType]PrattPrefixFn),
		infixFns:     make(map[token.TokenType]PrattInfixFn),
		postfixFns:   make(map[token.TokenType]PrattPostfixFn),
	}
}

func (p *PrattParser) RegisterPrefixFn(tokenType token.TokenType, fn PrattPrefixFn) {
	p.prefixFns[tokenType] = fn
}

func (p *PrattParser) RegisterInfixFn(tokenType token.TokenType, fn PrattInfixFn) {
	p.infixFns[tokenType] = fn
}

func (p *PrattParser) RegisterPostfixFn(tokenType token.TokenType, fn PrattPostfixFn) {
	p.postfixFns[tokenType] = fn
}

// func (p *PrattParser) ParseExpression(left ast.Node) ast.Node {
