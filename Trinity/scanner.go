package trinity

type Scanner struct {
	Source  string // The Source string
	Tokens  []Token
	start   int
	current int
	line    int
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

	scanner.Tokens = append(scanner.Tokens, Token{ // Adding and End of File TOken at end
		TokenType: EOF,
		Line:      scanner.line,
		Lexeme:    "EOF",
		Literal:   nil,
	})
}
