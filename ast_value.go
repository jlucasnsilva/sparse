package sparse

type (
	// FloatNode ...
	FloatNode struct {
		Row   int
		Col   int
		Value float64
	}

	// IntNode ...
	IntNode struct {
		Row   int
		Col   int
		Value uint64
	}

	// IdentifierNode ...
	IdentifierNode struct {
		Row   int
		Col   int
		Value string
	}
)

// Pos ...
func (n *FloatNode) Pos() (int, int) {
	return n.Row, n.Col
}

// Pos ...
func (n *IntNode) Pos() (int, int) {
	return n.Row, n.Col
}

// Pos ...
func (n *IdentifierNode) Pos() (int, int) {
	return n.Row, n.Col
}
