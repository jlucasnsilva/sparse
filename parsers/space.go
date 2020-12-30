package parsers

import (
	"errors"
	"fmt"
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

func parseBlank(value string) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	return &Blank{Value: len(value)}, nil
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

// ValueString ...
func (n *Blank) ValueString() string {
	return fmt.Sprintf("[blank:%v]", n.Value)
}

// ParseNewline ...
func ParseNewline(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValue(s, parseNewline)
}

func parseNewline(r rune) (sparse.Node, error) {
	if r != '\n' {
		return nil, errors.New("Not a newline")
	}
	return &Newline{}, nil
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

// ValueString ...
func (n *Newline) ValueString() string {
	return "\\n"
}

// ParseSpace ...
func ParseSpace(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, unicode.IsSpace, parseSpace)
}

func parseSpace(value string) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not white space")
	}
	return &Space{Value: len(value)}, nil
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

// ValueString ...
func (n *Space) ValueString() string {
	return "\\n"
}
