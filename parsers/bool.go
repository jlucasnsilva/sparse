package parsers

import "github.com/jlucasnsilva/sparse"

type (
	// BoolNode ...
	BoolNode struct {
		Row   int
		Col   int
		Value bool
	}
)

// Bool ...
func Bool(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := sparse.Or(
		ThisWord("true"),
		ThisWord("false"),
	)

	r, node, err := p(s)
	if err != nil {
		return s, nil, err
	}

	word := node.(*WordNode)
	value := true
	if word.Value == "false" {
		value = false
	}

	row, col := s.Position()
	result := &BoolNode{
		Row:   row,
		Col:   col,
		Value: value,
	}
	return r, result, nil
}

// Position ...
func (n *BoolNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *BoolNode) Equals(m sparse.Node) bool {
	v, ok := m.(*BoolNode)
	return ok && v.Value == n.Value
}

// String ...
func (n *BoolNode) String() string {
	return toString("BoolNode", n.Row, n.Col, n.Value)
}
