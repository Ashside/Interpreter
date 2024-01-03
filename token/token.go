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

	ASSIGN   = "="  // ASSIGN 赋值
	PLUS     = "+"  // PLUS 加号
	MINUS    = "-"  // MINUS 减号
	BANG     = "!"  // BANG 感叹号
	ASTERISK = "*"  // ASTERISK 星号
	SLASH    = "/"  // SLASH 斜杠
	LT       = "<"  // LT 小于号
	GT       = ">"  // GT 大于号
	EQ       = "==" // EQ 等于号
	NOT_EQ   = "!=" // NOT_EQ 不等于号

	COMMA     = "," // COMMA 逗号
	SEMICOLON = ";" // SEMICOLON 分号
	LPAREN    = "(" // LPAREN 左括号
	RPAREN    = ")" // RPAREN 右括号
	LBRACE    = "{" // LBRACE 左花括号
	RBRACE    = "}" // RBRACE 右花括号

	FUNCTION = "FUNCTION" // FUNCTION 函数
	LET      = "LET"      // LET 标识量
	IF       = "IF"       // IF
	ELSE     = "ELSE"     // ELSE
	RETURN   = "RETURN"   // RETURN
	TRUE     = "TRUE"     // TRUE
	FALSE    = "FALSE"    // FALSE
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// LookupIdent 判断标识符是否是关键字，关键字包括let、fn等
func LookupIdent(ident string) TokenType {
	// 如果是关键字，就返回关键字对应的token类型
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// 否则就返回标识符
	return IDENT

}
