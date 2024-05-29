package reader

type TokenType string

const (
	KEYWORD = "KEYWORD"
	SYMBOL  = "SYMBOL"
	CHAR    = "CHAR"
	EOL     = "EOL"
	EOT     = "EOT"
	UNKNOWN = "UNKNOWN"
)

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(t TokenType, s string) Token {
	return Token{Type: t, Literal: s}
}
