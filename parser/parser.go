package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"Interp/token"
)

type Parser struct {
	l         *lexer.Lexer //指向lexer
	curToken  token.Token  //当前token
	peekToken token.Token  //下一个token
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	//读取两个token，将curToken和peekToken都初始化
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() //读取下一个token

}

func (p *Parser) ParseProgram() *ast.Program {
	return nil

}
