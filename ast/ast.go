package ast

import (
	"Interp/token"
	"bytes"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}
type Statement interface {
	Node
	statementNode()
}
type Expression interface {
	Node
	expressionNode()
}

// Program 根节点
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

// String 返回当前program中所有语句的字符串，实际工作委托给statement的String方法
func (p *Program) String() string {
	var out bytes.Buffer //声明一个bytes.Buffer
	//遍历p.Statements，将每个语句的字符串拼接起来
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String 返回当前let语句的字符串
func (ls *LetStatement) String() string {
	var out bytes.Buffer //声明一个bytes.Buffer
	//拼接字符串
	out.WriteString(ls.TokenLiteral() + " ") //let
	out.WriteString(ls.Name.String())        //标识符
	out.WriteString(" = ")                   // =
	if ls.Value != nil {
		out.WriteString(ls.Value.String()) //表达式
	}
	out.WriteString(";") //分号
	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
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

// String 返回当前return语句的字符串
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer //声明一个bytes.Buffer
	//拼接字符串
	out.WriteString(rs.TokenLiteral() + " ") //return
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String()) //返回值
	}
	out.WriteString(";") //分号
	return out.String()
}

// ExpressionStatement 表达式语句
type ExpressionStatement struct {
	Token      token.Token // 表达式的第一个token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {

}

// TokenLiteral 返回当前token的字面量，实现了Node接口
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String 返回当前表达式语句的字符串
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {

}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {

}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer //声明一个bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {

}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *InfixExpression) String() string {
	var out bytes.Buffer //声明一个bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {

}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

type BlockStatement struct {
	Token      token.Token // { 词法单元
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {

}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {

}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())
	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {

}
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token //(词法单元
	Function  Expression  // 标识符或函数字面量
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {

}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	var args []string
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()

}
