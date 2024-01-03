package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"Interp/token"
	"fmt"
)

type Parser struct {
	l         *lexer.Lexer //指向lexer
	curToken  token.Token  //当前token
	peekToken token.Token  //下一个token
	errors    []string     //错误信息
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}} //初始化一个parser
	//读取两个token，将curToken和peekToken都初始化
	p.nextToken()
	p.nextToken()
	return p
}
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError 当peekToken不是期望的token时，将错误信息添加到p.errors
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	//将错误信息添加到p.errors
	p.errors = append(p.errors, msg)
}
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken() //读取下一个token

}

// ParseProgram parseStatement 解析语句
// NOTE: 作为一个成员函数，通过遍历p.curToken.Type来解析语句，但是只返回根节点
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}              //初始化一个program根节点
	program.Statements = []ast.Statement{} //初始化program.Statements，将其置为空

	// 循环解析语句，直到遇到token.EOF
	for !p.curTokenIs(token.EOF) {
		// 解析语句
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt) //将stmt添加到program.Statements
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET: //如果当前token是let，就解析let语句
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()

	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	//初始化一个let语句
	stmt := &ast.LetStatement{Token: p.curToken}

	//判断下一个token是否是标识符
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	//初始化一个标识符
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//判断下一个token是否是等号
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	//TODO: 跳过表达式，直到遇到分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseReturnStatement 构造一个ast.ReturnStatement，并将当前词法单元放置到Token字段中
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	//初始化一个return语句
	stmt := &ast.ReturnStatement{Token: p.curToken}
	p.nextToken()
	//TODO: 跳过表达式，直到遇到分号
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()

	}
	return stmt

}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek 判断下一个token是否是期望的token t，如果是就读取下一个token
func (p *Parser) expectPeek(t token.TokenType) bool {
	//如果下一个token是期望的token，就读取下一个token
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		//如果下一个token不是期望的token，就将错误信息添加到p.errors
		p.peekError(t)
		return false
	}
}
