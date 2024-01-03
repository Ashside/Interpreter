package ast

import "Interp/token"

type Node interface {
	TokenLiteral() string
}
type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

// TokenLiteral 如果当前program中有语句，就返回第一个语句的字面量
func (p *Program) TokenLiteral() string {
	// 如果当前program中有语句，就返回第一个语句的字面量
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral() //返回第一个语句的字面量
	} else {
		return ""
	}
}

// LetStatement let语句
type LetStatement struct {
	Token token.Token // token.LET 标识符
	Name  *Identifier // 标识符
	Value Expression  // 表达式
}

// statementNode 表示这是一个语句，实现了Statement接口
func (ls *LetStatement) statementNode() {

}

// TokenLiteral 返回当前token的字面量，实现了Node接口
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier 标识符
type Identifier struct {
	Token token.Token // token.IDENT 标识符
	Value string      // 标识符的值
}

// expressionNode 表示这是一个表达式，实现了Expression接口
func (i *Identifier) expressionNode() {

}

// TokenLiteral 返回当前token的字面量，实现了Node接口
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement return语句
type ReturnStatement struct {
	Token       token.Token // token.RETURN 标识符
	ReturnValue Expression  // 返回值
}

// statementNode 表示这是一个语句，实现了Statement接口
func (rs *ReturnStatement) statementNode() {

}

// TokenLiteral 返回当前token的字面量，实现了Node接口
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
