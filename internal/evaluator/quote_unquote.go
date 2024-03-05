package evaluator

import (
	"github.com/Ryoga-exe/monkey/internal/ast"
	"github.com/Ryoga-exe/monkey/internal/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
