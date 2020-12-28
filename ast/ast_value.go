package ast

type (
	// Float ...
	Float struct {
		Row   int
		Col   int
		Value float64
	}

	// Int ...
	Int struct {
		Row   int
		Col   int
		Value uint64
	}

	// Identifier ...
	Identifier struct {
		Row   int
		Col   int
		Value string
	}

	// Rune ...
	Rune struct {
		Row   int
		Col   int
		Value rune
	}
)

// Position ...
func (n *Float) Position() (int, int) {
	return n.Row, n.Col
}

// Child ...
func (n *Float) Child(i int) Node {
	panic("Nodes of type 'Float' don't have children")
}

// Position ...
func (n *Int) Position() (int, int) {
	return n.Row, n.Col
}

// Child ...
func (n *Int) Child(i int) Node {
	panic("Nodes of type 'Int' don't have children")
}

// Position ...
func (n *Identifier) Position() (int, int) {
	return n.Row, n.Col
}

// Child ...
func (n *Identifier) Child(i int) Node {
	panic("Nodes of type 'Identifier' don't have children")
}

// Position ...
func (n *Rune) Position() (int, int) {
	return n.Row, n.Col
}

// Child ...
func (n *Rune) Child(i int) Node {
	panic("Nodes of type 'Rune' don't have children")
}
