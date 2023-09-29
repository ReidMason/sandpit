package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	THERES  = "theres"
	THIS    = "this"
	MOVIE   = "movie"
	CALLED  = "called"
	WHICH   = "which"
	IS      = "is"
	INT     = "INT"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	IDENT   = "IDENT"
)

var keywords = map[string]TokenType{
	"theres": THERES,
	"this":   THIS,
	"movie":  MOVIE,
	"called": CALLED,
	"which":  WHICH,
	"is":     IS,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
