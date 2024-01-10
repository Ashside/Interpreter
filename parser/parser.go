package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"Interp/token"
	"fmt"
	"strconv"
)

type Parser struct {
	l      *lexer.Lexer //指向lexer
	errors []string     //错误信息

	curToken  token.Token //当前token
	peekToken token.Token //下一个token

	prefixParseFns map[token.TokenType]prefixParseFn //前缀解析函数
	infixParseFns  map[token.TokenType]infixParseFn  //中缀解析函数
}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)

)

// precedences 定义了每个运算符的优先级，表示了运算符右约束能力。
// 当右约束能力达到最大值，那么当前解析的结果，即分配给leftExp的值就不会传递给下一个运算符关联的infixParseFn
// 也就是说，leftExp不会成为左子节点，因为此时parseExpression函数中for循环的条件为false
var precedences = map[token.TokenType]int{
	token.EQ:       EQUALS,
	token.NOT_EQ:   EQUALS,
	token.LT:       LESSGREATER,
	token.GT:       LESSGREATER,
	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.SLASH:    PRODUCT,
	token.ASTERISK: PRODUCT,
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}} //初始化一个parser
	//读取两个token，将curToken和peekToken都初始化
	p.nextToken()
	p.nextToken()
	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
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
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	//初始化一个let语句
	stmt := &ast.LetStatement{Token: p.curToken}

	//判断下一个token是否是标识符
	if !p.expectPeekMove(token.IDENT) {
		return nil
	}

	//初始化一个标识符
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	//判断下一个token是否是等号
	if !p.expectPeekMove(token.ASSIGN) {
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

func (p *Parser) parseExpressionStatement() ast.Statement {
	defer untrace(trace("parseExpressionStatement: " + p.curToken.Literal))
	stmt := &ast.ExpressionStatement{Token: p.curToken} //初始化一个表达式语句
	stmt.Expression = p.parseExpression(LOWEST)         //解析表达式,LOWEST表示最低优先级
	//如果下一个token是分号，就读取下一个token
	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

// parseExpression 解析表达式
// NOTE: version 1 检查前缀位置受否有与p.curToken.Type对应的解析函数，如果有就调用该函数并返回，否则返回nil
func (p *Parser) parseExpression(precedence int) ast.Expression {
	defer untrace(trace("parseExpression: " + p.curToken.Literal))
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()
	// 该方法尝试为下一个词法单元查找infixParseFns，如果找到了这个函数，就用prefixParseFn返回的表达式作为参数调用这个函数。
	// 循环重复执行，直到遇见优先级更高的词法单元为止
	// 在前一次读入操作符，这一次读到数字的时候，将不会进入循环
	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken() // 因为在infix中需要用到p.curToken，所以这里先移动词法单元
		leftExp = infix(leftExp)
	}
	return leftExp
}

// parseIdentifier 关联解析函数，将当前词法单元及其字面量提供给*ast.Identifier的Token和Value字段，并返回该节点
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.curToken,
		Value: p.curToken.Literal,
	}
}

// parseIntegerLiteral 调用了strconv.ParseInt，将p.curToken的字面值赋给Expression(IntegerLiteral)的value字段
func (p *Parser) parseIntegerLiteral() ast.Expression {
	defer untrace(trace("parseIntegerLiteral: " + p.curToken.Literal))
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

// parsePrefixExpression 会调用p.nextToken()来前移词法单元，开始的时候p.curToken是前缀运算符，返回时指向前缀表达式的操作数
func (p *Parser) parsePrefixExpression() ast.Expression {
	defer untrace(trace("parsePrefixExpression: " + p.curToken.Literal))
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Right:    nil,
	}
	p.nextToken()                                //移动词法单元
	expression.Right = p.parseExpression(PREFIX) //按照前缀优先级解析
	return expression
}

// parseInfixExpression 会调用p.nextToken()来前移词法单元，开始的时候p.curToken是中缀运算符，返回时指向中缀表达式的右操作数
// 参数left是中缀表达式的左操作数
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	defer untrace(trace("parseInfixExpression: " + p.curToken.Literal))
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}
	precedence := p.curPrecedence()                  //获取当前优先级
	p.nextToken()                                    //移动词法单元，此时p.curToken指向操作符下一个词法单元
	expression.Right = p.parseExpression(precedence) //按照当前优先级解析，递归调用parseExpression
	return expression
}

// curTokenIs 判断当前token是否是期望的token
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs 判断下一个token是否是期望的token
func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeekMove 判断下一个token是否是期望的token t，如果是就读取下一个token
func (p *Parser) expectPeekMove(t token.TokenType) bool {
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

// prefixParseFns 前缀解析函数
type prefixParseFn func() ast.Expression

// infixParseFns 中缀解析函数
type infixParseFn func(ast.Expression) ast.Expression

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
func (p *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)

}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST // 如果precedences没有对应的优先级，就使用默认的最低优先级
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST // 如果precedences没有对应的优先级，就使用默认的最低优先级
}

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: p.curToken,
		Value: p.curTokenIs(token.TRUE),
	}
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	exp := p.parseExpression(LOWEST)
	if !p.expectPeekMove(token.RPAREN) {
		return nil
	}
	return exp
}
