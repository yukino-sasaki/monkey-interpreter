package evaluator

import (
	"monkey-interpreter/01/monkey/ast"
	"monkey-interpreter/01/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}