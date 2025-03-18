package parser

type Parser struct {
	Tokens  []any // SWITCH: Token
	current int
}

func NewParser(tokens []any) Parser { // SWITCH: Token
	return Parser{
		Tokens:  tokens,
		current: 0,
	}
}
