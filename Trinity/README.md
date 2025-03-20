# Trinity 

The official Lexer for the Neo Programming Language.


## Scanner

The Scanner provides a struct that goes through all the source code for neo, which will be for a file with format .neo and creates an array of tokens for it.

### scanner.go
this file contains the Scanner struct declaration along with the `NewScanner()`
function to create a new Scanner instance.

In this file there is also the `ScanTokens()` function which iterates over the neo source code string.
```go
func (scanner *Scanner) ScanTokens() {
	for {
		if scanner.isAtEnd() {
			break
		}

		// current is the current character we are at, and start is where we started
		// this is for multiple character long Tokens
		scanner.start = scanner.current
		scanner.ScanToken()
	}

	scanner.Tokens = append(scanner.Tokens, lextoken.Token{ // Adding and End of File TOken at end
		TokenType: lextoken.EOF,
		Line:      scanner.line,
		Lexeme:    "EOF",
		Literal:   nil,
	})
}
```

here we loop over our source code till we reach the end of the string. For every character we iterate over we execute the `ScanToken()` that scans the token and stores the Token parsed inside the Tokens attribute for our Scanner.

The reason we have `scanner.start = scanner.current` is because our current attribute represents our current index in the source code string and start represents the start of our current lexeme/token, for example a `!=` token spans multiple characters and we need to be able to know when it started.


### scannerUtils.go

This file contains several helper functions we use across the Lexer, the functions are:-

- `addToken()`
- `isAtEnd()`
- `nextChar()`
- `isDigit()`
- `isAlpha()`

### scanTokens.go

In this file we solely have the `ScanToken()` function which converts the current lexeme/string to a Token. Its a big collection of switch case.

The reason we only check the single/two character in `ScanToken()` using switch case and handle Keywords, and Identifiers (variable names etc) is that they may overlap. An Identifier 'orchid' may overlap with our keyword 'or' hence we first parse the full string and then check if it matches a keyword in our grammar. 

### literalHandlers.go

In this file we have several handlers that handle the situation for different Literals such as integer/string or handling Identifier/Keywords.

It contains:- 

- `handleString()`
- `handleNumber()`
- `haneleIdentifier()`

## Token

this package contains an implementation of the Token enumeration and also the keyword hashmap and the Token struct implementation.




