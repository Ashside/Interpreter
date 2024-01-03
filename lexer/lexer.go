package lexer

import "Interp/token"

type Lexer struct {
	input        string
	position     int  //当前字符的位置
	readPosition int  //当前读取字符的位置，即当前字符的下一个字符的位置
	ch           byte //当前字符
}

// New 创建一个新的Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input} //初始化一个Lexer l，将input赋值给l.input
	l.readChar()              //初始化l.ch、l.position、l.readPosition
	return l
}

// readChar 读取下一个字符，将l.readPosition向后移动一位
func (l *Lexer) readChar() {
	//如果读取到了input的末尾，就将ch设置为0，表示到达了文件末尾
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken 读取下一个token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token //声明一个token

	l.skipWhitespace() //跳过空白符

	//根据当前字符，判断当前token的类型
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		//EOF
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		//如果当前字符是字母或者下划线，就读取标识符
		if isLetter(l.ch) {
			tok.Literal = l.readCharIdent()
			tok.Type = token.LookupIdent(tok.Literal) //判断标识符是否是关键字
			return tok                                //返回一个标识符token
		} else if isDigit(l.ch) {
			tok.Literal = l.readCharIdent()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}

	}

	l.readChar()
	return tok
}

// newToken 根据传入的token类型和字符创建一个新的token
// 例如：newToken(token.ASSIGN, '=')，将返回一个类型为token.ASSIGN，值为'='的token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter 判断ch是否是字母或者下划线，这里用于判断标识符。决定了解释器中的标识符只能由字母或者下划线组成，也可修改已满足其他需求
func isLetter(ch byte) bool {
	//如果是字母或者下划线，就返回true
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// readIdentifier 读取标识符，这里的标识符指的是let、fn等
func (l *Lexer) readIdentifier() string {
	//记录当前字符的位置
	currentPosition := l.position
	//如果当前字符是字母或者下划线，就一直读取下一个字符
	for isLetter(l.ch) {
		l.readChar()
	}
	//返回从current_position到l.position的字符串
	return l.input[currentPosition:l.position]
}

// skipWhitespace 跳过空白符
func (l *Lexer) skipWhitespace() {
	//如果当前字符是空格、换行符、回车符、制表符，就一直读取下一个字符
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// isDigit 判断ch是否是数字
func isDigit(ch byte) bool {
	//如果是数字，就返回true
	return '0' <= ch && ch <= '9'
}

// readNumber 读取数字
func (l *Lexer) readNumber() string {
	// 记录当前字符的位置
	currentPosition := l.position
	// 如果当前字符是数字，就一直读取下一个字符
	for isDigit(l.ch) {
		l.readChar()
	}
	// 返回从current_position到l.position的字符串
	return l.input[currentPosition:l.position]
}

// @TODO: 完全可以整合number 和 identifier到一个函数readCharIdent中
// @DONE: 整合完毕
// readCharIdent 读取标识符和数字
func (l *Lexer) readCharIdent() string {
	// 记录当前字符的位置
	currentPosition := l.position
	// 如果当前字符是数字或者字母，就一直读取下一个字符
	for isDigit(l.ch) || isLetter(l.ch) {
		l.readChar()
	}
	// 返回从current_position到l.position的字符串
	return l.input[currentPosition:l.position]
}
