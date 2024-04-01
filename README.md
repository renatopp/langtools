# LANGTOOLS

This is a my collection of tools I use to develop my toy programming languages.

## Installation

```bash
go get github.com/renatopp/langtools
```

## Content

- `runes` - utility package that provide rune checkers such as `isWhitespace`.
- `tokens` - contains types for `Char` and `Token`, representing the lexical info from source code.
- `lexers` - contains all utilities for tokenization. Mostly based on the `BaseLexer` class.
- `parsers` - contains all utilities for parsing tokens into ASTs. Mostly based on the `BaseParser` class.
- `asts` - contains the types for AST `Node`.

Check the [./examples](./examples) folder for detailed usage.

## Usage

