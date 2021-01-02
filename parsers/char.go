package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// CharNode ...
	CharNode struct {
		Row   int
		Col   int
		Value rune
	}
)

// Char ...
func Char(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	var node sparse.Node
	singleQuoteParser := ThisRune('\'')
	r, _, err := singleQuoteParser(s)
	if err != nil {
		return r, nil, err
	}
	r, node, err = parseValue(r, createChar)
	if err != nil {
		return r, nil, err
	}
	r, _, err = singleQuoteParser(r)
	if err != nil {
		return r, nil, err
	}
	return r, node, nil
}

func createChar(r rune, row, col int) (sparse.Node, error) {
	return &CharNode{Value: r, Row: row, Col: col}, nil
}

// Position ...
func (n *CharNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *CharNode) Equals(m sparse.Node) bool {
	v, ok := m.(*CharNode)
	return ok && v.Value == n.Value
}

// String ...
func (n *CharNode) String() string {
	return toString("CharNode", n.Row, n.Col, fmt.Sprintf("'%c'", n.Value))
}
