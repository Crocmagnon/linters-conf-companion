package fatcontext

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "fatcontext",
	Doc:      "detects nested contexts in loops",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspctr := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.RangeStmt)(nil),
	}

	inspctr.Preorder(nodeFilter, func(node ast.Node) {
		var body *ast.BlockStmt

		switch typedNode := node.(type) {
		case *ast.RangeStmt:
			body = typedNode.Body
		default:
			return
		}

		processBody(pass, body)
	})

	return nil, nil
}

func processBody(pass *analysis.Pass, body *ast.BlockStmt) {

}
