package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	t.Log("TestLetStatements Called")

	input := `
return 5;
return 10;
return 993322;
`
	//初始化一个lexer
	l := lexer.NewLexer(input)
	//初始化一个parser
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	// 解析let语句，所以只能解析3元素的program
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got=%q", returnStmt.TokenLiteral())
		}
	}
}

// testLetStatement 测试let语句
func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	t.Log("testLetStatement Called")

	//判断当前语句是否是let语句
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got=%q", s.TokenLiteral())
		return false
	}

	//将s转换成*ast.LetStatement类型
	letStmt, ok := s.(*ast.LetStatement) //类型断言
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	//判断letStmt.Name.Value是否等于name
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
		return false
	}

	//判断letStmt.Name.TokenLiteral()是否等于name
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

// checkParserErrors 检查parser的错误，如果有错误，就打印错误信息
func checkParserErrors(t *testing.T, p *Parser) {
	//如果p.errors的长度为0，表示没有错误，直接返回
	if len(p.errors) == 0 {
		return
	}
	//否则就打印错误信息
	t.Errorf("parser has %d errors", len(p.errors))
	//遍历p.errors，打印错误信息
	for _, msg := range p.errors {
		t.Errorf("parser error: %q", msg)
	}
	//测试失败
	t.FailNow()
}

func TestIdentifierExpression(t *testing.T) {
	t.Log("TestIdentifierExpression Called")
	input := "foobar;"
	//初始化一个lexer
	l := lexer.NewLexer(input)
	//初始化一个parser
	//首先用语法分析器检查是否有错误
	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	//解析的program只能有一条语句
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements, got=%d", len(program.Statements))
	}

	//检查这一条唯一的语句是否为*ast.ExpressionStatement类型
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	//检查这条语句的Expression是否为*ast.Identifier类型
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier, got=%T", stmt.Expression)
	}

	//判断ident.Value是否等于foobar
	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s, got=%s", "foobar", ident.Value)
	}

	//判断ident.TokenLiteral()是否等于foobar
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s, got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.NewLexer(input)

	p := NewParser(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	//解析的program只能有一条语句
	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements, got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral, got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d, got=%d", 5, literal.Value)
	}
	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral not %s, got=%s", "5", literal.TokenLiteral())
	}

}
