package trinity

//TODO: Number Literals
//TODO: identifiers
//TODO: I did it wrong, keywords will be handled seperately

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {

	//Brackets
	case '{':
		scanner.addToken(BRACE_LEFT, nil)
		break
	case '}':
		scanner.addToken(BRACE_RIGHT, nil)
		break
	case '(':
		scanner.addToken(PAREN_LEFT, nil)
		break
	case ')':
		scanner.addToken(PAREN_RIGHT, nil)
		break

		//Arithmetic
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(FUNC_RETURN, nil)
			break
		}
		scanner.addToken(MINUS, nil)
		break
	case '+':
		scanner.addToken(PLUS, nil)
		break
	case '*':
		scanner.addToken(MULTIPLY, nil)
		break
	case '/':
		if scanner.nextChar('/') {
			for !scanner.nextChar('\n') && !scanner.isAtEnd() {
				scanner.current++
			}
			scanner.line++
		} else {
			scanner.addToken(DIVIDE, nil)
		}
		break

		// Comparison
	case '>':
		if scanner.nextChar('=') {
			scanner.addToken(GREATER_EQ, nil)
			break
		}
		scanner.addToken(GREATER, nil)
		break
	case '<':
		if scanner.nextChar('=') {
			scanner.addToken(LESS_EQ, nil)
			break
		}
		scanner.addToken(LESS, nil)
		break
	case '=':
		if scanner.nextChar('=') {
			scanner.addToken(EQ_EQ, nil)
			break
		}
		scanner.addToken(EQUAL, nil)
		break
		//Logical
	case '!':
		if scanner.nextChar('=') {
			scanner.addToken(NOT_EQ, nil)
			break
		}
		scanner.addToken(NOT, nil)
		break
	case '&':
		if scanner.nextChar('&') {
			scanner.addToken(AND, nil)
		}
		break
	case '|':
		if scanner.nextChar('|') {
			scanner.addToken(OR, nil)
		}
		break

	// Single Char (whats left)
	case ',':
		scanner.addToken(COMMA, nil)
		break
	case ':':
		scanner.addToken(COLON, nil)
		break
	case ';':
		scanner.addToken(SEMI_COLON, nil)
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
		break
	}
}
