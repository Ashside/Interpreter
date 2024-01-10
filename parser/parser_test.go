package parser

import (
	"Interp/ast"
	"Interp/lexer"
	"fmt"
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

/*
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
}*/

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

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    int64
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)
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

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		{
			if !ok {
				t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
			}
			if exp.Operator != tt.operator {
				t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
			}
			if !testIntegerLiteral(t, exp.Right, tt.value) {
				return
			}
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value interface{}) bool {

	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}
	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}
	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}
	return true
}

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}
func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	//将exp转换成*ast.Boolean类型
	boolean, ok := exp.(*ast.Boolean)
	if !ok {
		t.Errorf("exp not *ast.Boolean, got=%T", exp)
		return false
	}
	//判断boolean.Value是否等于value
	if boolean.Value != value {
		t.Errorf("boolean.Value not %t, got=%t", value, boolean.Value)
		return false
	}
	//判断boolean.TokenLiteral()是否等于value
	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("boolean.TokenLiteral not %t, got=%s", value, boolean.TokenLiteral())
		return false
	}
	return true
}
func testInfixExpression(t *testing.T, exp ast.Expression, left interface{}, operator string, right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfixExpression. got=%T", exp)
		return false
	}
	//判断左边的值是否正确
	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}
	//判断运算符是否正确
	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}
	//判断右边的值是否正确
	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}
	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)
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
		// 判断stmt.Expression是否为*ast.InfixExpression类型
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.InfixExpression. got=%T", stmt.Expression)
		}

		//判断左边的值是否正确
		if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
			return
		}
		//判断运算符是否正确
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}
		//判断右边的值是否正确
		if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((-a) * b)"},
		{"!-a", "(!(-a))"},
		{"a + b + c", "((a + b) + c)"},
		{"a + b - c", "((a + b) - c)"},
		{"a * b * c", "((a * b) * c)"},
		{"a * b / c", "((a * b) / c)"},
		{"a + b / c", "(a + (b / c))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)"},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))"},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))"},
		{"true", "true"},
		{"false", "false"},
		{"3 > 5 == false", "((3 > 5) == false)"},
		{"3 < 5 == true", "((3 < 5) == true)"},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(-(5 + 5))"},
		{"!(true == true)", "(!(true == true))"},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input           string
		expectedBoolean bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
		p := NewParser(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program has not enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		boolean, ok := stmt.Expression.(*ast.Boolean)
		if !ok {
			t.Fatalf("exp not *ast.Boolean. got=%T", stmt.Expression)
		}
		if boolean.Value != tt.expectedBoolean {
			t.Errorf("boolean.Value not %t. got=%t", tt.expectedBoolean,
				boolean.Value)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := `
if (x < y) { x }
`
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
	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt is not ast.IfExpression. got=%T", stmt.Expression)
	}
	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}
	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}
	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}
	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `
if (x < y) { x } else { y }
`
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
	exp, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("stmt is not ast.IfExpression. got=%T", stmt.Expression)
	}
	if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
		return
	}
	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d\n", len(exp.Consequence.Statements))
	}
	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.Consequence.Statements[0])
	}
	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}
	if exp.Alternative == nil {
		t.Errorf("exp.Alternative.Statements was nil. got=%+v", exp.Alternative)
	}
	if len(exp.Alternative.Statements) != 1 {
		t.Errorf("exp.Alternative.Statements was not 1 statements. got=%+v", len(exp.Alternative.Statements))
	}
	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statements[0] is not ast.ExpressionStatement. got=%T", exp.Alternative.Statements[0])
	}
	if !testIdentifier(t, alternative.Expression, "y") {
		return
	}
}

func TestFunctionLiteralParsing(t *testing.T) {
	input := `fn(x, y) { x + y; }`
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
	// 检查stmt.Expression是否为*ast.FunctionLiteral类型
	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("stmt is not ast.FunctionLiteral. got=%T", stmt.Expression)
	}
	// 检查参数列表是否正确
	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}
	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")
	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements. got=%d\n", len(function.Body.Statements))
	}
	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}
	testInfixExpression(t, bodyStmt.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{"fn(){}", []string{}},
		{"fn(x){}", []string{"x"}},
		{"fn(x,y,z){}", []string{"x", "y", "z"}},
	}
	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)
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
		// 检查stmt.Expression是否为*ast.FunctionLiteral类型
		function, ok := stmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("stmt is not ast.FunctionLiteral. got=%T", stmt.Expression)
		}
		// 检查参数列表是否正确
		if len(function.Parameters) != len(tt.expectedParams) {
			t.Fatalf("length parameters wrong. want %d, got=%d\n", len(tt.expectedParams), len(function.Parameters))
		}
		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}
