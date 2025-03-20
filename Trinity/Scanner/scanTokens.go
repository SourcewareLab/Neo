package lexscanner

import token "NeoLang/Token"

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {
	//Brackets
	case '{':
		scanner.addToken(token.BRACE_LEFT, nil)

	case '}':
		scanner.addToken(token.BRACE_RIGHT, nil)

	case '(':
		scanner.addToken(token.PAREN_LEFT, nil)

	case ')':
		scanner.addToken(token.PAREN_RIGHT, nil)

		//Arithmetic
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(token.FUNC_RETURN, nil)
			break
		}
		scanner.addToken(token.MINUS, nil)

	case '+':
		scanner.addToken(token.PLUS, nil)

	case '*':
		scanner.addToken(token.MULTIPLY, nil)

	case '/':
		if scanner.nextChar('/') {
			for !scanner.nextChar('\n') && !scanner.isAtEnd() {
				scanner.current++
			}
			scanner.line++
		} else {
			scanner.addToken(token.DIVIDE, nil)
		}

		// Comparison
	case '>':
		if scanner.nextChar('=') {
			scanner.addToken(token.GREATER_EQ, nil)
			break
		}
		scanner.addToken(token.GREATER, nil)

	case '<':
		if scanner.nextChar('=') {
			scanner.addToken(token.LESS_EQ, nil)
			break
		}
		scanner.addToken(token.LESS, nil)

	case '=':
		if scanner.nextChar('=') {
			scanner.addToken(token.EQ_EQ, nil)

		}
		scanner.addToken(token.EQUAL, nil)

		//Logical
	case '!':
		if scanner.nextChar('=') {
			scanner.addToken(token.NOT_EQ, nil)
			break
		}
		scanner.addToken(token.NOT, nil)

	case '&':
		if scanner.nextChar('&') {
			scanner.addToken(token.AND, nil)
		}

	case '|':
		if scanner.nextChar('|') {
			scanner.addToken(token.OR, nil)
		}
		//throw err

	// Single Char (whats left)
	case ',':
		scanner.addToken(token.COMMA, nil)

	case ':':
		scanner.addToken(token.COLON, nil)

	case ';':
		scanner.addToken(token.SEMI_COLON, nil)

	case '\'': // '' and "" both are string
	case '"':
		scanner.handleString()

	case ' ': // Cases for whitespace or stuff we dont need, add actual tokens above this
	case '\t':
	case '\r':

	case '\n':
		scanner.line++

	default: // Handling of number/float literals and identifiers (var_name, func_name etc) TBD here
		if scanner.isDigit(char) {
			scanner.handleNumber()
		} else if scanner.isAlpha(char) {
			scanner.handleIdentifier()
		}
	}
}
