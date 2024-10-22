package token

// We could make the type be represented by an int or byte for performance,
// but I am not very concerned in performance as this won't be production lang
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"

	//	identifiers and literals
	IDENT = "IDENT"
	INT	  = "INT"

	// operators
	PLUS	 = "+"
	ASSIGN 	 = "="
	MINUS 	 = "-"
	BANG 	 = "!"
	SLASH	 = "/"
	ASTERISK = "*"

	LT 	   = "<"
	GT 	   = ">"
	EQ 	   = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET		 = "LET"
	TRUE	 = "TRUE"
	FALSE	 = "FALSE"
	IF 		 = "IF"
	ELSE	 = "ELSE"
	RETURN	 = "RETURN"
)

type Token struct {
	Type TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"true":TRUE,
	"false":FALSE,
	"if":IF,
	"else":ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}