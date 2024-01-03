package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
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
