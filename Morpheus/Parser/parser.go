package parser

import token "NeoLang/Token"

type Parser struct {
	Tokens  []token.Type
	current int
}

func NewParser(tokens []token.Type) Parser {
	return Parser{
		Tokens:  tokens,
		current: 0,
	}
}
