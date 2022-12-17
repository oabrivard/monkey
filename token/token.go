package token

type TokenType byte

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = iota
	EOF
	// Identifiers + literals
	IDENT // add, foobar, x, y, ...
	INT   // 1343456
	// Operators
	ASSIGN
	PLUS
	// Delimiters
	COMMA
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE
	// Keywords
	FUNCTION
	LET
)

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INT:
		return "INT"
	case ASSIGN:
		return "return"
	case PLUS:
		return "+"
	case COMMA:
		return ","
	case SEMICOLON:
		return ";"
	case LPAREN:
		return "("
	case RPAREN:
		return ")"
	case LBRACE:
		return "{"
	case RBRACE:
		return "}"
	case FUNCTION:
		return "FUNCTION"
	case LET:
		return "LET"
	}

	panic("unknown token type")
}

func LookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"fn":  FUNCTION,
		"let": LET,
	}

	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}