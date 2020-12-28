package ast

import (
	"fmt"
)

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

	// Word ...
	Word struct {
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

	// Blank ...
	Blank struct {
		Row   int
		Col   int
		Value int // length
	}

	// Newline ...
	Newline struct {
		Row int
		Col int
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

// ValueString ...
func (n *Float) ValueString() string {
	return fmt.Sprint(n.Value)
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

// ValueString ...
func (n *Int) ValueString() string {
	return fmt.Sprint(n.Value)
}

// Position ...
func (n *Word) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Word) Equals(m Node) bool {
	v, ok := m.(*Word)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Word) Child(i int) Node {
	panic("Nodes of type 'Word' don't have children")
}

// Children ...
func (n *Word) Children() int {
	panic("Nodes of type 'Word' don't have children")
}

// ValueString ...
func (n *Word) ValueString() string {
	return n.Value
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

// ValueString ...
func (n *Rune) ValueString() string {
	return fmt.Sprintf("'%c'", n.Value)
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

// ValueString ...
func (n *String) ValueString() string {
	return fmt.Sprintf("%c%v%c", n.Bracket, n.Value, n.Bracket)
}

// Equals ...
func (n *Blank) Equals(m Node) bool {
	v, ok := m.(*Blank)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Blank) Child(i int) Node {
	panic("Nodes of type 'Blank' don't have children")
}

// Children ...
func (n *Blank) Children() int {
	panic("Nodes of type 'Blank' don't have children")
}

// Position ...
func (n *Blank) Position() (int, int) {
	return n.Row, n.Col
}

// ValueString ...
func (n *Blank) ValueString() string {
	return fmt.Sprintf("[blank:%v]", n.Value)
}

// Equals ...
func (n *Newline) Equals(m Node) bool {
	_, ok := m.(*Newline)
	return ok
}

// Child ...
func (n *Newline) Child(i int) Node {
	panic("Nodes of type 'Newline' don't have children")
}

// Children ...
func (n *Newline) Children() int {
	panic("Nodes of type 'Newline' don't have children")
}

// Position ...
func (n *Newline) Position() (int, int) {
	return n.Row, n.Col
}

// ValueString ...
func (n *Newline) ValueString() string {
	return "\\n"
}
