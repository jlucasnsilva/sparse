package sparseutil

import (
	"github.com/jlucasnsilva/sparse"
	"github.com/jlucasnsilva/sparse/ast"
)

// SingleQuoteString ...
func SingleQuoteString(s sparse.Scanner) (sparse.Scanner, ast.Node, error) {
	p := sparse.OneString('\'')
	return p(s)
}

// DoubleQuoteString ...
func DoubleQuoteString(s sparse.Scanner) (sparse.Scanner, ast.Node, error) {
	p := sparse.OneString('"')
	return p(s)
}

// BackTickString ...
func BackTickString(s sparse.Scanner) (sparse.Scanner, ast.Node, error) {
	p := sparse.OneString('`')
	return p(s)
}

// Char ...
func Char(s sparse.Scanner) (sparse.Scanner, ast.Node, error) {
	var node ast.Node
	r, _, err := sparse.ThisRune('\'')(s)
	if err != nil {
		return r, nil, err
	}
	r, node, err = sparse.Rune(r)
	if err != nil {
		return r, nil, err
	}
	r, _, err = sparse.ThisRune('\'')(r)
	if err != nil {
		return r, nil, err
	}
	return r, node, nil
}
