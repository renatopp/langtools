package asts

import "github.com/renatopp/langtools/tokens"

type Node interface {
	GetToken() tokens.Token
	String() string
	Children() []Node
	Traverse(int, func(int, Node))
}
