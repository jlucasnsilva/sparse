package parsers

import (
	"errors"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
	// BlankNode ...
	BlankNode struct {
		Row   int
		Col   int
		Value int // length
	}

	// NewlineNode ...
	NewlineNode struct {
		Row int
		Col int
	}

	// SpaceNode ...
	SpaceNode struct {
		Row   int
		Col   int
		Value int // length
	}
)

// Blank ...
func Blank(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, isBlank, createSpace)
}

func createSpace(value string, row, col int) (sparse.Node, error) {
	return createSpaceNode(value, row, col, false)
}

// DismissBlank ...
func DismissBlank(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, isBlank, dismissSpace)
}

func dismissSpace(value string, row, col int) (sparse.Node, error) {
	return createSpaceNode(value, row, col, true)
}

func createSpaceNode(value string, row, col int, dismiss bool) (sparse.Node, error) {
	var result *BlankNode
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	if !dismiss {
		result = &BlankNode{
			Value: len(value),
			Row:   row,
			Col:   col,
		}
	}
	return result, nil
}

func isBlank(r rune) bool {
	return unicode.IsSpace(r) && r != '\n'
}

// Equals ...
func (n *BlankNode) Equals(m sparse.Node) bool {
	v, ok := m.(*BlankNode)
	return ok && v.Value == n.Value
}

// Child ...
func (n *BlankNode) Child(i int) sparse.Node {
	panic("Nodes of type 'BlankNode' don't have children")
}

// Children ...
func (n *BlankNode) Children() int {
	panic("Nodes of type 'BlankNode' don't have children")
}

// Position ...
func (n *BlankNode) Position() (int, int) {
	return n.Row, n.Col
}

// String ...
func (n *BlankNode) String() string {
	return toString("BlankNode", n.Row, n.Col, n.Value)
}

// Newline ...
func Newline(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValue(s, createNewline)
}

func createNewline(r rune, row, col int) (sparse.Node, error) {
	if r != '\n' {
		return nil, errors.New("Not a newline")
	}
	return &NewlineNode{Row: row, Col: col}, nil
}

// DismissNewline ...
func DismissNewline(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, isNewline, dismissSpace)
}

func isNewline(r rune) bool {
	return r == '\n'
}

// Equals ...
func (n *NewlineNode) Equals(m sparse.Node) bool {
	_, ok := m.(*NewlineNode)
	return ok
}

// Child ...
func (n *NewlineNode) Child(i int) sparse.Node {
	panic("Nodes of type 'NewlineNode' don't have children")
}

// Children ...
func (n *NewlineNode) Children() int {
	panic("Nodes of type 'NewlineNode' don't have children")
}

// Position ...
func (n *NewlineNode) Position() (int, int) {
	return n.Row, n.Col
}

// String ...
func (n *NewlineNode) String() string {
	return toString("NewlineNode", n.Row, n.Col, "\\n")
}

// Space ...
func Space(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, unicode.IsSpace, parseSpace)
}

// DismissSpace ...
func DismissSpace(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, isBlank, dismissSpace)
}

func parseSpace(value string, row, col int) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	result := &SpaceNode{
		Value: len(value),
		Row:   row,
		Col:   col,
	}
	return result, nil
}

// Equals ...
func (n *SpaceNode) Equals(m sparse.Node) bool {
	_, ok := m.(*SpaceNode)
	return ok
}

// Child ...
func (n *SpaceNode) Child(i int) sparse.Node {
	panic("Nodes of type 'NewlineNode' don't have children")
}

// Children ...
func (n *SpaceNode) Children() int {
	panic("Nodes of type 'NewlineNode' don't have children")
}

// Position ...
func (n *SpaceNode) Position() (int, int) {
	return n.Row, n.Col
}

// String ...
func (n *SpaceNode) String() string {
	return toString("SpaceNode", n.Row, n.Col, n.Value)
}
