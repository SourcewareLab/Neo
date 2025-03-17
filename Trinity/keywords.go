package trinity

var keywords = map[string]TokenType{
	"int":      INT,
	"float":    FLOAT,
	"string":   STRING,
	"function": FUNC_TYPE,
	"null":     NULL,
	"bool":     BOOL,
	"err":      ERROR_TYPE,
	"mut":      MUT,
	"const":    CONST,
	"if":       IF,
	"else":     ELSE,
	"elseif":   ELSE_IF,
	"for":      FOR,
	"inf":      INF,
	"range":    RANGE,
	"println":  PRINT,
	"error":    ERROR_WORD,
	"return":   FUNC_RETURN,
	"fn":       FUNC_WORD,
}
