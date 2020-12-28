package sparse

import "github.com/jlucasnsilva/sparse/ast"

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node ast.Node, err error)
)
