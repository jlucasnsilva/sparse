package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// Char ...
	Char struct {
		Row   int
		Col   int
		Value rune
	}
)

// ParseChar ...
func ParseChar(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	var node sparse.Node
	singleQuoteParser := ParseThisRune('\'')
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
	return &Char{Value: r, Row: row, Col: col}, nil
}

// Position ...
func (n *Char) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Char) Equals(m sparse.Node) bool {
	v, ok := m.(*Char)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Char) Child(i int) sparse.Node {
	panic("Nodes of type 'Char' don't have children")
}

// Children ...
func (n *Char) Children() int {
	panic("Nodes of type 'Char' don't have children")
}

// String ...
func (n *Char) String() string {
	return toString("Char", n.Row, n.Col, fmt.Sprintf("'%c'", n.Value))
}
