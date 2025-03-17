package trinity

func (scanner *Scanner) nextChar(char byte) bool {
	if scanner.isAtEnd() || scanner.Source[scanner.current] != char {
		return false
	}

	scanner.current++
	return true
}

func (scanner *Scanner) addToken(tokenType TokenType, literal any) {
	lexeme := scanner.Source[scanner.start:scanner.current]

	scanner.Tokens = append(scanner.Tokens, Token{
		TokenType: tokenType,
		Literal:   literal,
		Lexeme:    lexeme,
		Line:      scanner.line,
	})
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.Source)
}

func (scanner *Scanner) matchString(str string) bool { // should throw error later when we implement syntax errors
	for _, val := range str {
		if scanner.Source[scanner.current] != byte(val) {
			return false
		}

		scanner.current++
	}

	return true
}

func (scanner *Scanner) handleString() {
	for {
		if scanner.nextChar('"') || scanner.isAtEnd() {
			break
		}
		scanner.current++
	}

	if scanner.isAtEnd() {
	} // should throw error

	// scanner.current++ // for the last "

	literal := scanner.Source[scanner.start+1 : scanner.current-1]

	scanner.Tokens = append(scanner.Tokens, Token{
		TokenType: DUB_QUOTES,
		Literal:   literal,
		Line:      scanner.line,
		Lexeme:    "\"\"",
	})
}
