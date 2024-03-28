# LANGTOOLS

This is a collection of tools used develop programming languages.

```go
input := []byte("Hello, World!")
lexer := langtools.NewLexer(input, tokenizers.Word)
for i, t := lexer.Iter() {
	fmt.Printf("[%03d] %d:%d (%s) %s\n", i, token.Line, token.Column, token.Type, token.Literal)
}
```