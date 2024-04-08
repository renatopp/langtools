package asts

import (
	"fmt"
	"strings"
)

type TraverseFn func(depth int, node Node)

// Print the AST nodes recursively.
func Print(node Node, ident string) {
	Traverse(node, func(depth int, node Node) {
		if node == nil {
			fmt.Printf("%s%s\n", strings.Repeat(ident, depth), "nil")
			return
		}

		fmt.Printf("%s%s\n", strings.Repeat(ident, depth), node.String())
	})
}

// Traverse the AST nodes recursively.
func Traverse(node Node, f TraverseFn) {
	traverse(0, node, f)
}

func traverse(depth int, node Node, f TraverseFn) {
	f(depth, node)

	if node == nil {
		return
	}

	for _, child := range node.Children() {
		traverse(depth+1, child, f)
	}
}
