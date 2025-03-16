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
	})
}

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {
	case '{':
		scanner.addToken(BRACE_LEFT, NULL)
		break
	case '}':
		scanner.addToken(BRACE_RIGHT, NULL)
		break
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(FUNC_RETURN, NULL) // Its null for now, it should have the dataType as a literal
		}
		break
	case 'f':
		if scanner.nextChar('n') {
			scanner.addToken(FUNC_WORD, NULL)
		} // should give error else
		break
	case ' ': // Cases for whitespace or stuff we dont need, add actual tokens above this
	case '\t':
	case '\r':
		break
	default:
		scanner.addToken(ERR, NULL)
		break
	}
}

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
