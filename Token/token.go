package token

type Token struct {
	TokenType Type
	Lexeme    string // The actual token string is called lexeme
	Literal   any    // A Lexeme for a literal value, this has to be interpreted seperate from normal lexemes
	Line      int    // The current line in the program
}
