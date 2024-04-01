package asts

import "github.com/renatopp/langtools/tokens"

type Node interface {
	GetToken() tokens.Token
	String() string
	Children() []Node
	Traverse(int, func(int, Node))
}

type BaseNode struct {
	Token tokens.Token
}

func (n *BaseNode) GetToken() tokens.Token {
	return n.Token
}

func (n *BaseNode) String() string {
	return string(n.Token.Type) + " " + n.Token.Literal
}

func (n *BaseNode) Children() []Node {
	return []Node{}
}

func (n *BaseNode) Traverse(level int, f func(int, Node)) {
	f(level, n)
	for _, child := range n.Children() {
		child.Traverse(level+1, f)
	}
}
