package sparse

import (
	"errors"
	"fmt"

	"github.com/jlucasnsilva/sparse/ast"
)

// And ...
func And(parsers ...ParserFunc) ExprParserFunc {
	return func(s Scanner) (Scanner, []ast.Node, error) {
		var (
			nodes = make([]ast.Node, 0, 10)
			node  ast.Node
			err   error
			r     = s
		)
		for i, p := range parsers {
			r, node, err = p(r)
			if err != nil {
				return s, nil, fmt.Errorf("invalid parser at %v: %v", i, err)
			}
			nodes = append(nodes, node)
		}
		return r, nodes, nil
	}
}

// Or ...
func Or(parsers ...ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, ast.Node, error) {
		for _, p := range parsers {
			if r, node, err := p(s); err == nil {
				return r, node, nil
			}
		}
		return s, nil, errors.New("No matching parser")
	}
}

// Some ...
func Some(target ParserFunc, separator ParserFunc) ExprParserFunc {
	return func(s Scanner) (Scanner, []ast.Node, error) {
		var (
			nodes = make([]ast.Node, 0, 10)
			node  ast.Node
			err   error
			r     = s
			t     = r
		)
		for err == nil {
			if t, node, err = target(r); err == nil {
				nodes = append(nodes, node)
				r = t
			}
		}
		if len(nodes) < 1 {
			return s, nil, errors.New("not a single match")
		}
		return r, nodes, nil
	}
}
