# Changelog

<!-- NEWER -->

## v0.2.7 (2024-04-08)

- Traverse now checks for nil nodes

## v0.2.6 (2024-04-07)

- Changing ast.Node interface to use token pointer

## v0.2.5 (2024-04-07)

- Improving debug string in tokens
- Fixing EatString and EatRawString error positions

## v0.2.4 (2024-04-07)

- Fixing error interface (Range func)

## v0.2.3 (2024-04-07)

- Octal now accepting 0o and 0O formats

## v0.2.2 (2024-04-07)

- Numbers are now eaten with _ character

## v0.2.1 (2024-04-07)

- Fixing lexer tests
- BaseLexer.EatIndetifierWith call signature

## v0.2.0 (2024-04-07)

- Tokens now save the complete range (fromLine, fromColumn, toLine, toColumn)
- Complete refactor on Token to adapt these information
- Tokens now are used as pointers
- Adding utility functions to handle errors messages

## v0.1.2 (2024-04-02)

- Adding error interface

## v0.1.1 (2024-04-02)

- Adding PrevToken() to BaseLexer.
- No * methods in Token and Char.
- Removing Traverse from asts.Node interface.
- Removing BaseNode from asts.
- Adding Traverse function to ast package.

## v0.1.0 (2024-04-01)

- Breaking refactor for organization
- Merging eaters into BaseLexer
- Adding BaseParser and PrattParser
- Adding Mathematical example

## v0.0.0 (2024-03-28)

- Initial package with basic lexer
