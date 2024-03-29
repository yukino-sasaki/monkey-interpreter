package ast

import (
	"monkey-interpreter/01/monkey/token"
	"testing"
)


func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal : "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "let"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		//　なぜq？？
		t.Errorf("program.String() wrong got %q", program.String())
	}
}