package token

type TokenType string

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "ASSIGN"
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

var keyword = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keyword[ident]; ok {
		return tok
	}
	return IDENT
}

type Token struct {
	Type    TokenType
	Literal string
}
