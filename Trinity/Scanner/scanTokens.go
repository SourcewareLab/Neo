package lexscanner

import lextoken "NeoLang/Trinity/Token"

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {

	//Brackets
	case '{':
		scanner.addToken(lextoken.BRACE_LEFT, nil)
		break
	case '}':
		scanner.addToken(lextoken.BRACE_RIGHT, nil)
		break
	case '(':
		scanner.addToken(lextoken.PAREN_LEFT, nil)
		break
	case ')':
		scanner.addToken(lextoken.PAREN_RIGHT, nil)
		break

		//Arithmetic
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(lextoken.FUNC_RETURN, nil)
			break
		}
		scanner.addToken(lextoken.MINUS, nil)
		break
	case '+':
		scanner.addToken(lextoken.PLUS, nil)
		break
	case '*':
		scanner.addToken(lextoken.MULTIPLY, nil)
		break
	case '/':
		if scanner.nextChar('/') {
			for !scanner.nextChar('\n') && !scanner.isAtEnd() {
				scanner.current++
			}
			scanner.line++
		} else {
			scanner.addToken(lextoken.DIVIDE, nil)
		}
		break

		// Comparison
	case '>':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.GREATER_EQ, nil)
			break
		}
		scanner.addToken(lextoken.GREATER, nil)
		break
	case '<':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.LESS_EQ, nil)
			break
		}
		scanner.addToken(lextoken.LESS, nil)
		break
	case '=':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.EQ_EQ, nil)
			break
		}
		scanner.addToken(lextoken.EQUAL, nil)
		break
		//Logical
	case '!':
		if scanner.nextChar('=') {
			scanner.addToken(lextoken.NOT_EQ, nil)
			break
		}
		scanner.addToken(lextoken.NOT, nil)
		break
	case '&':
		if scanner.nextChar('&') {
			scanner.addToken(lextoken.AND, nil)
		}
		break
	case '|':
		if scanner.nextChar('|') {
			scanner.addToken(lextoken.OR, nil)
		}
		break

	// Single Char (whats left)
	case ',':
		scanner.addToken(lextoken.COMMA, nil)
		break
	case ':':
		scanner.addToken(lextoken.COLON, nil)
		break
	case ';':
		scanner.addToken(lextoken.SEMI_COLON, nil)
		break
	case '\'': // '' and "" both are string
	case '"':
		scanner.handleString()
		break
	case ' ': // Cases for whitespace or stuff we dont need, add actual tokens above this
	case '\t':
	case '\r':
		break
	case '\n':
		scanner.line++
		break
	default: // Handling of number/float literals and identifiers (var_name, func_name etc) TBD here
		c := scanner.Source[scanner.current]
		if scanner.isDigit(c) {
			scanner.handleNumber()
			break
		} else if scanner.isAlpha(c) {
			scanner.handleIdentifier()
		}

		break
	}
}
