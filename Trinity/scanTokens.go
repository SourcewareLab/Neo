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

		// Keywords
	case 'f':
		if scanner.nextChar('n') {
			scanner.addToken(FUNC_WORD, nil)
		} else if scanner.matchString("loat") {
			scanner.addToken(FLOAT, nil)
		} else if scanner.matchString("unction") {
			scanner.addToken(FUNC_TYPE, nil)
		} // should give error else
		break
	case 'n':
		if scanner.matchString("ull") {
			scanner.addToken(NULL, nil)
		}
		break
	case 'b':
		if scanner.matchString("ool") {
			scanner.addToken(BOOL, nil)
		}
		break
	case 'e':
		if scanner.matchString("rr") {
			scanner.addToken(ERROR_TYPE, nil)
		} else if scanner.matchString("lse") {
			if scanner.matchString(" if") { // nested if for Optimization, les go
				scanner.addToken(ELSE_IF, nil)
			} else {
				scanner.addToken(ELSE, nil)
			}
		}
		break
	case 'm':
		if scanner.matchString("ain()") {
			scanner.addToken(MAIN_FUNC, nil)
		} else if scanner.matchString("ut") {
			scanner.addToken(MUT, nil)
		}
		break
	case 'c':
		if scanner.matchString("onst") {
			scanner.addToken(CONST, nil)
		}
		break
	case 'i':
		if scanner.nextChar('n') {
			if scanner.nextChar('f') {
				scanner.addToken(INF, nil)
			} else if scanner.nextChar('t') {
				scanner.addToken(INT, nil)
			}
		} else if scanner.nextChar('f') {
			scanner.addToken(IF, nil)
		}
		break
	case 'r':
		if scanner.matchString("ange") {
			scanner.addToken(RANGE, nil)
		} else if scanner.matchString("eturn") {
			scanner.addToken(FUNC_RETURN, nil)
		}
		break
	case 'p':
		if scanner.matchString("rintln") {
			scanner.addToken(PRINT, nil)
		}
		break
	case 's':
		if scanner.matchString("tring") {
			scanner.addToken(STRING, nil)
		}
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
