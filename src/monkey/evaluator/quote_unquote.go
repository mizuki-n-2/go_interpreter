package evaluator

import (
	"github.com/mizuki-n-2/go_interpreter/src/monkey/ast"
	"github.com/mizuki-n-2/go_interpreter/src/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
