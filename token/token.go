package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAl = "ILLEGAL"
	EOF     = "EOF"

	//标识量、字面量
	IDENT = "IDENT"
	INT   = "INT"

	//运算符
	ASSIGN = "="
	PLUS   = "+"

	//分隔符
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
