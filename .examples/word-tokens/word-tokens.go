package main

import (
	"fmt"
	"io/ioutil"

	"github.com/renatopp/langtools"
	"github.com/renatopp/langtools/tokenizers"
)

func main() {
	input, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	lexer := langtools.NewLexer(input, tokenizers.Word)
	for i, token := range lexer.Iter() {
		fmt.Printf("[%03d] %d:%d (%s) %s\n", i, token.Line, token.Column, token.Type, token.Literal)
	}

	if lexer.HasErrors() {
		fmt.Println("Lexer errors:")
		for _, e := range lexer.Errors() {
			fmt.Printf("[%d:%d] %s\n", e.Line, e.Column, e.Error())
		}
	}
}
