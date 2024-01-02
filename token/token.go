package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // ILLEGAL 表示非法字符
	EOF     = "EOF"     // EOF 表示到达文件末尾

	IDENT = "IDENT" // IDENT 标识量
	INT   = "INT"   // INT 整型

	ASSIGN = "=" // ASSIGN 赋值
	PLUS   = "+" // PLUS 加号

	COMMA     = "," // COMMA 逗号
	SEMICOLON = ";" // SEMICOLON 分号
	LPAREN    = "(" // LPAREN 左括号
	RPAREN    = ")" // RPAREN 右括号
	LBRACE    = "{" // LBRACE 左花括号
	RBRACE    = "}" // RBRACE 右花括号

	FUNCTION = "FUNCTION" // FUNCTION 函数
	LET      = "LET"      // LET 标识量

)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	// 如果是关键字，就返回关键字对应的token类型
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// 否则就返回标识符
	return IDENT

}
