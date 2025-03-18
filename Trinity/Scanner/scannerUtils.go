package lexscanner

import lextoken "NeoLang/Trinity/Token"

func (scanner *Scanner) nextChar(char byte) bool {
	if scanner.isAtEnd() || scanner.Source[scanner.current] != char {
		return false
	}

	scanner.current++
	return true
}

func (scanner *Scanner) addToken(tokenType lextoken.TokenType, literal any) {
	lexeme := scanner.Source[scanner.start:scanner.current]

	scanner.Tokens = append(scanner.Tokens, lextoken.Token{
		TokenType: tokenType,
		Literal:   literal,
		Lexeme:    lexeme,
		Line:      scanner.line,
	})
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.Source)
}

func (scanner *Scanner) isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func (scanner *Scanner) isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

// func (scanner *scanner) matchstring(str string) bool { // should throw error later when we implement syntax errors
//	for _, val := range str {
//		if scanner.source[scanner.current] != byte(val) {
//			return false
//		}
//
//		scanner.current++
//	}
//
//	return true
// }
