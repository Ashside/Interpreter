package ast

import (
	"Interp/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"}, //token.LET
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"}, //token.IDENT
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"}, //token.IDENT
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong, got=%q", program.String())
	}
}
