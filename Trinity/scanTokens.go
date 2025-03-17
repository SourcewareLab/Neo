package trinity

func (scanner *Scanner) ScanToken() {
	char := scanner.Source[scanner.current]
	scanner.current++

	switch char {
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
	case '-':
		if scanner.nextChar('>') {
			scanner.addToken(FUNC_RETURN, nil)
		}
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
	case 'f':
		if scanner.nextChar('n') {
			scanner.addToken(FUNC_WORD, nil)
		} // should give error else
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
		if scanner.matchString("nt") {
			scanner.addToken(INT, nil)
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
	default:
		scanner.addToken(ERR, nil)
		break
	}
}
