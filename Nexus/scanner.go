package nexus

type Scanner struct {
	Source string // The Source string
	Tokens []Token
}

func newScanner(source string) Scanner {
	return Scanner{
		Source: source,
	}
}
