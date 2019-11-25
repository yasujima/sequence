package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	NODE = "NODE"
	IDENT = "IDENT"
	STRING = "STRING"
	CRLF = "CRLF"
	
	COLON = ":"
	BULLET = "-"
	LINE = "-"
	LT = "<"
	GT = ">"
	HASH = "#"

)

func LookupIdent(ident string) TokenType {
	return IDENT
}
	
	
	
	
