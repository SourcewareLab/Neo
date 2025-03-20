package lexscanner

import (
	token "NeoLang/Token"
	"strconv"
)

func (scanner *Scanner) handleString() {
	for {
		if scanner.nextChar('"') || scanner.isAtEnd() {
			break
		}
		scanner.current++
	}

	if scanner.isAtEnd() {
	} // should throw error

	literal := scanner.Source[scanner.start+1 : scanner.current-1]

	scanner.Tokens = append(scanner.Tokens, token.Token{
		TokenType: token.STRING,
		Literal:   literal,
		Line:      scanner.line,
		Lexeme:    "string",
	})
}

func (scanner *Scanner) handleNumber() {
	for {
		if !scanner.isDigit(scanner.Source[scanner.current]) || scanner.isAtEnd() {
			break
		}
		scanner.current++
	}

	if scanner.nextChar('.') && scanner.isDigit(scanner.Source[scanner.current]) {
		scanner.current++ // to use the '.'

		for {
			if !scanner.isDigit(scanner.Source[scanner.current]) || scanner.isAtEnd() {
				break
			}
			scanner.current++
		}

		s := scanner.Source[scanner.start:scanner.current]
		float, _ := strconv.ParseFloat(s, 64)

		scanner.Tokens = append(scanner.Tokens, token.Token{
			Lexeme:    "float",
			Literal:   float,
			TokenType: token.FLOAT,
			Line:      scanner.line,
		})
		return
	}

	s := scanner.Source[scanner.start:scanner.current]
	int, _ := strconv.ParseInt(s, 10, 64)

	scanner.Tokens = append(scanner.Tokens, token.Token{
		Lexeme:    "int",
		Literal:   int,
		TokenType: token.INT,
		Line:      scanner.line,
	})
}

func (scanner *Scanner) handleIdentifier() {
	for {
		if !scanner.isAlpha(scanner.Source[scanner.current]) || scanner.isAtEnd() {
			break
		}

		scanner.current++
	}

	lexeme := scanner.Source[scanner.start:scanner.current]

	val, found := token.Keywords[lexeme]
	if !found {
		scanner.addToken(token.IDENTIFIER, nil)
		return
	}

	scanner.addToken(val, nil)
}

func (scanner *Scanner) nextChar(char byte) bool {
	if scanner.isAtEnd() || scanner.Source[scanner.current] != char {
		return false
	}

	scanner.current++
	return true
}

func (scanner *Scanner) addToken(tokenType token.Type, literal any) {
	lexeme := scanner.Source[scanner.start:scanner.current]

	scanner.Tokens = append(scanner.Tokens, token.Token{
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
