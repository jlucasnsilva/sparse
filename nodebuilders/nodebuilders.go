package nodebuilders

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// ArrayNode ...
	ArrayNode struct {
		Row   int
		Col   int
		Value []sparse.Node
	}

	// Dismiss ...
	Dismiss struct{}

	// Array ...
	Array struct {
		nodes []sparse.Node
	}
)

// Build ...
func (b Dismiss) Build() sparse.Node {
	return nil
}

// Add ...
func (b Dismiss) Add(sparse.Node) {
	// Do nothing
}

// Reset ...
func (b Dismiss) Reset() {
	// Do nothing
}

// Build ...
func (b *Array) Build() sparse.Node {
	if len(b.nodes) < 1 {
		return nil
	}

	first := b.nodes[0]
	row, col := first.Position()
	return &ArrayNode{
		Value: b.nodes,
		Row:   row,
		Col:   col,
	}
}

// Add ...
func (b *Array) Add(n sparse.Node) {
	if n != nil {
		b.nodes = append(b.nodes, n)
	}
}

// Reset ...
func (b *Array) Reset() {
	b.nodes = nil
}

// String ...
func (n *ArrayNode) String() string {
	return fmt.Sprintf(
		"nodebuilders.ArrayNode{ Row: %v, Col: %v, Value «%v» }",
		n.Row, n.Col, n.Value,
	)
}

// Equals ...
func (n *ArrayNode) Equals(m sparse.Node) bool {
	mn, ok := m.(*ArrayNode)
	if !ok {
		return false
	}
	if len(mn.Value) != len(n.Value) {
		return false
	}

	narr := n.Value
	marr := mn.Value
	for i := 0; i < len(narr); i++ {
		if !narr[i].Equals(marr[i]) {
			return false
		}
	}
	return true
}

// Position ...
func (n *ArrayNode) Position() (int, int) {
	return n.Row, n.Col
}
