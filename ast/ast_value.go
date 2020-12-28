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

	// String ...
	String struct {
		Row     int
		Col     int
		Bracket rune
		Value   string
	}
)

// Position ...
func (n *Float) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Float) Equals(m Node) bool {
	v, ok := m.(*Float)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Float) Child(i int) Node {
	panic("Nodes of type 'Float' don't have children")
}

// Children ...
func (n *Float) Children() int {
	panic("Nodes of type 'Float' don't have children")
}

// Position ...
func (n *Int) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Int) Equals(m Node) bool {
	v, ok := m.(*Int)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Int) Child(i int) Node {
	panic("Nodes of type 'Int' don't have children")
}

// Children ...
func (n *Int) Children() int {
	panic("Nodes of type 'Int' don't have children")
}

// Position ...
func (n *Identifier) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Identifier) Equals(m Node) bool {
	v, ok := m.(*Identifier)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Identifier) Child(i int) Node {
	panic("Nodes of type 'Identifier' don't have children")
}

// Children ...
func (n *Identifier) Children() int {
	panic("Nodes of type 'Identifier' don't have children")
}

// Position ...
func (n *Rune) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Rune) Equals(m Node) bool {
	v, ok := m.(*Rune)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Rune) Child(i int) Node {
	panic("Nodes of type 'Rune' don't have children")
}

// Children ...
func (n *Rune) Children() int {
	panic("Nodes of type 'Rune' don't have children")
}

// Equals ...
func (n *String) Equals(m Node) bool {
	v, ok := m.(*String)
	return ok && v.Value == n.Value
}

// Child ...
func (n *String) Child(i int) Node {
	panic("Nodes of type 'String' don't have children")
}

// Children ...
func (n *String) Children() int {
	panic("Nodes of type 'String' don't have children")
}

// Position ...
func (n *String) Position() (int, int) {
	return n.Row, n.Col
}
