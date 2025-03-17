package trinity

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
	case '/':
		if scanner.nextChar('/') {
			for !scanner.nextChar('\n') && !scanner.isAtEnd() {
				scanner.current++
			}
			scanner.line++
		} else {
			scanner.addToken(DIVIDE, NULL)
		}
		break
	case 'f':
		if scanner.nextChar('n') {
			scanner.addToken(FUNC_WORD, NULL)
		} // should give error else
		break
	case 'm':
		if scanner.matchString("ain()") {
			scanner.addToken(MAIN_FUNC, NULL)
		}
		break
	case 'i':
		if scanner.matchString("nt") {
			scanner.addToken(INT, NULL)
		}
		break
	case 'p':
		if scanner.matchString("rintln") {
			scanner.addToken(PRINT, NULL)
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
		scanner.addToken(ERR, NULL)
		break
	}
}
