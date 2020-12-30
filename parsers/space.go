package parsers

import (
	"errors"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
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

	// Space ...
	Space struct {
		Row   int
		Col   int
		Value int // length
	}
)

// ParseBlank ...
func ParseBlank(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, isBlank, parseBlank)
}

func parseBlank(value string, row, col int) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	result := &Blank{
		Value: len(value),
		Row:   row,
		Col:   col,
	}
	return result, nil
}

func isBlank(r rune) bool {
	return unicode.IsSpace(r) && r != '\n'
}

// Equals ...
func (n *Blank) Equals(m sparse.Node) bool {
	v, ok := m.(*Blank)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Blank) Child(i int) sparse.Node {
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

// String ...
func (n *Blank) String() string {
	return toString("Blank", n.Row, n.Col, n.Value)
}

// ParseNewline ...
func ParseNewline(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValue(s, parseNewline)
}

func parseNewline(r rune, row, col int) (sparse.Node, error) {
	if r != '\n' {
		return nil, errors.New("Not a newline")
	}
	return &Newline{Row: row, Col: col}, nil
}

// Equals ...
func (n *Newline) Equals(m sparse.Node) bool {
	_, ok := m.(*Newline)
	return ok
}

// Child ...
func (n *Newline) Child(i int) sparse.Node {
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

// String ...
func (n *Newline) String() string {
	return toString("Newline", n.Row, n.Col, "\\n")
}

// ParseSpace ...
func ParseSpace(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, unicode.IsSpace, parseSpace)
}

func parseSpace(value string, row, col int) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	result := &Space{
		Value: len(value),
		Row:   row,
		Col:   col,
	}
	return result, nil
}

// Equals ...
func (n *Space) Equals(m sparse.Node) bool {
	_, ok := m.(*Space)
	return ok
}

// Child ...
func (n *Space) Child(i int) sparse.Node {
	panic("Nodes of type 'Newline' don't have children")
}

// Children ...
func (n *Space) Children() int {
	panic("Nodes of type 'Newline' don't have children")
}

// Position ...
func (n *Space) Position() (int, int) {
	return n.Row, n.Col
}

// String ...
func (n *Space) String() string {
	return toString("Space", n.Row, n.Col, n.Value)
}
