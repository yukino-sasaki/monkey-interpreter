package ast

import (
	"reflect"
	"testing"
)

type ModifierFunc func(Node) Node

func TestModify(t *testing.T) {
	one := func() Expression { return &IntegerLiteral{Value: 1}}
	two := func() Expression { return &IntegerLiteral{Value: 2}}

	turnOneIntoTwo := func(node Node) Node {
		integer, ok := node.(*IntegerLiteral)
		if !ok {
			return node
		}

		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	tests := []struct {
		input Node
		expected Node
	} {
		{
			one(),
			two(),
		},
		{
			&Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: one()},
				},
			},
			&Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: two()},
				},
			},
		},
	}

	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)

		equal := reflect.DeepEqual(modified, tt.expected)
		if !equal {
			t.Errorf("not equal. got%#v, want%#v", modified, tt.expected)
		}
	}
}

func Modify(node Node, modifier ModifierFunc) Node{
	switch node := node.(type) {
	case *Program:
		for i, statement := range node.Statements {
			node.Statements[i], _ = Modify(statement, modifier).(Statement)
		}
	case *ExpressionStatement:
		node.Expression = Modify(node.Expression, modifier).(Expression)
	}
	return modifier(node)
}