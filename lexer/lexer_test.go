package lexer

import (
	"Interp/token"
	"testing"
)

/*
func TestNextToken(t *testing.T) {
	input := `=+(){},;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},    //赋值
		{token.PLUS, "+"},      //加号
		{token.LPAREN, "("},    //左括号
		{token.RPAREN, ")"},    //右括号
		{token.LBRACE, "{"},    //左花括号
		{token.RBRACE, "}"},    //右花括号
		{token.COMMA, ","},     //逗号
		{token.SEMICOLON, ";"}, //分号
		{token.EOF, ""},
	}

	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}*/

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x,y){
	x + y;
};
let result = add(five,ten);
!-/*5;
5 < 10 > 5;

if (5 < 10){
	return true;
}else{
	return false;
}
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"}, //let
		{token.IDENT, "five"},
		{token.ASSIGN, "="},    //=
		{token.INT, "5"},       //5
		{token.SEMICOLON, ";"}, //;
		{token.LET, "let"},     //let
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},    //=
		{token.INT, "10"},      //10
		{token.SEMICOLON, ";"}, //;
		{token.LET, "let"},     //let
		{token.IDENT, "add"},
		{token.ASSIGN, "="},    //=
		{token.FUNCTION, "fn"}, //fn
		{token.LPAREN, "("},    //(
		{token.IDENT, "x"},
		{token.COMMA, ","}, //,
		{token.IDENT, "y"},
		{token.RPAREN, ")"}, //)
		{token.LBRACE, "{"}, //{
		{token.IDENT, "x"},
		{token.PLUS, "+"}, //+
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"}, //;
		{token.RBRACE, "}"},    //}
		{token.SEMICOLON, ";"}, //;
		{token.LET, "let"},     //let
		{token.IDENT, "result"},
		{token.ASSIGN, "="}, //=
		{token.IDENT, "add"},
		{token.LPAREN, "("}, //(
		{token.IDENT, "five"},
		{token.COMMA, ","}, //,
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},      //)
		{token.SEMICOLON, ";"},   //;
		{token.BANG, "!"},        //!
		{token.MINUS, "-"},       //-
		{token.SLASH, "/"},       ///
		{token.ASTERISK, "*"},    //*
		{token.INT, "5"},         //5
		{token.SEMICOLON, ";"},   //;
		{token.INT, "5"},         //5
		{token.LT, "<"},          //<
		{token.INT, "10"},        //10
		{token.GT, ">"},          //>
		{token.INT, "5"},         //5
		{token.SEMICOLON, ";"},   //;
		{token.IF, "if"},         //if
		{token.LPAREN, "("},      //(
		{token.INT, "5"},         //5
		{token.LT, "<"},          //<
		{token.INT, "10"},        //10
		{token.RPAREN, ")"},      //)
		{token.LBRACE, "{"},      //{
		{token.RETURN, "return"}, //return
		{token.TRUE, "true"},     //true
		{token.SEMICOLON, ";"},   //;
		{token.RBRACE, "}"},      //}
		{token.ELSE, "else"},     //else
		{token.LBRACE, "{"},      //{
		{token.RETURN, "return"}, //return
		{token.FALSE, "false"},   //false
		{token.SEMICOLON, ";"},   //;
		{token.RBRACE, "}"},      //}
		{token.EOF, ""},
	}
	l := NewLexer(input)
	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
