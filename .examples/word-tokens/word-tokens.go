package main

import (
	"fmt"
	"io/ioutil"

	"github.com/renatopp/langtools/lexers"
)

func main() {
	input, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		panic(err)
	}

	lexer := lexers.NewBaseLexer(input, lexers.WordTokenizer)
	i := 0
	for {
		token, eof := lexer.Next()
		if eof {
			break
		}

		fmt.Printf("[%03d] %d:%d (%s) %s\n", i, token.Line, token.Column, token.Type, token.Literal)
		i++
	}

	if lexer.HasErrors() {
		fmt.Println("Lexer errors:")
		for _, e := range lexer.Errors() {
			fmt.Printf("[%d:%d] %s\n", e.Line, e.Column, e.Error())
		}
	}
}
