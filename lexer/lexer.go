package lexer

import "Interp/token"

type Lexer struct {
	input        string
	position     int  //当前字符的位置
	readPosition int  //当前读取字符的位置，即当前字符的下一个字符的位置
	ch           byte //当前字符
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/**
 * 读取input中的下一个字符
 */
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
	var tok token.Token
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
	}
	l.readChar()
	return tok
}

/**
 * 生成一个新的token
 */
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
