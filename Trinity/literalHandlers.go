package trinity

import "strconv"

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

	scanner.Tokens = append(scanner.Tokens, Token{
		TokenType: STRING,
		Literal:   literal,
		Line:      scanner.line,
		Lexeme:    "\"\"",
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

		scanner.Tokens = append(scanner.Tokens, Token{
			Lexeme:    "float",
			Literal:   float,
			TokenType: FLOAT,
			Line:      scanner.line,
		})
		return
	}

	s := scanner.Source[scanner.start:scanner.current]
	int, _ := strconv.ParseInt(s, 10, 64)

	scanner.Tokens = append(scanner.Tokens, Token{
		Lexeme:    "int",
		Literal:   int,
		TokenType: INT,
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

	val, found := keywords[lexeme]
	if !found {
		scanner.addToken(IDENTIFIER, nil)
		return
	}

	scanner.addToken(val, nil)
}
