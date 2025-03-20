package lexscanner

import lextoken "NeoLang/Trinity/Token"

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {
	//Brackets
	case '{':
		scanner.addToken(lextoken.BRACE_LEFT, nil)

	case '}':
		scanner.addToken(lextoken.BRACE_RIGHT, nil)

	case '(':
		scanner.addToken(lextoken.PAREN_LEFT, nil)

	case ')':
		scanner.addToken(lextoken.PAREN_RIGHT, nil)

		//Arithmetic
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(lextoken.FUNC_RETURN, nil)
			break
		}
		scanner.addToken(lextoken.MINUS, nil)

	case '+':
		scanner.addToken(lextoken.PLUS, nil)

	case '*':
		scanner.addToken(lextoken.MULTIPLY, nil)

	case '/':
		if scanner.nextChar('/') {
			for !scanner.nextChar('\n') && !scanner.isAtEnd() {
				scanner.current++
			}
			scanner.line++
		} else {
			scanner.addToken(lextoken.DIVIDE, nil)
		}

		// Comparison
	case '>':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.GREATER_EQ, nil)
			break
		}
		scanner.addToken(lextoken.GREATER, nil)

	case '<':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.LESS_EQ, nil)
			break
		}
		scanner.addToken(lextoken.LESS, nil)

	case '=':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.EQ_EQ, nil)

		}
		scanner.addToken(lextoken.EQUAL, nil)

		//Logical
	case '!':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.NOT_EQ, nil)
			break
		}
		scanner.addToken(lextoken.NOT, nil)

	case '&':
		if scanner.nextChar('&') {
			scanner.addToken(lextoken.AND, nil)
		}

	case '|':
		if scanner.nextChar('|') {
			scanner.addToken(lextoken.OR, nil)
		}
		//throw err

	// Single Char (whats left)
	case ',':
		scanner.addToken(lextoken.COMMA, nil)

	case ':':
		scanner.addToken(lextoken.COLON, nil)

	case ';':
		scanner.addToken(lextoken.SEMI_COLON, nil)

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
