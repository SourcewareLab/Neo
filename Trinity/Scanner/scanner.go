package lexscanner

import (
	token "NeoLang/Token"
)

type Scanner struct {
	Source  string // The Source string
	Tokens  []token.Token
	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:  source,
		Tokens:  make([]token.Token, 0),
		start:   0,
		current: 0,
		line:    0,
	}
}

func (scanner *Scanner) ScanTokens() {
	for {
		if scanner.isAtEnd() {
			break
		}

		// current is the current character we are at, and start is where we started
		// this is for multiple character long Tokens
		scanner.start = scanner.current
		scanner.ScanToken()
	}

	scanner.Tokens = append(scanner.Tokens, token.Token{ // Adding and End of File TOken at end
		TokenType: token.EOF,
		Line:      scanner.line,
		Lexeme:    "EOF",
		Literal:   nil,
	})
}
