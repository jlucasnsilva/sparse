package parsers

import (
	"errors"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
	// WordNode ...
	WordNode struct {
		Row   int
		Col   int
		Value string
	}
)

// Word ...
func Word(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	count := 0
	check := func(r rune) bool {
		if count == 0 {
			count++
			return isWordFirst(r)
		}
		return isWord(r)
	}
	return parseValueWithWhile(s, check, createWord)
}

func createWord(value string, row, col int) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not an word")
	}
	return &WordNode{Value: value, Row: row, Col: col}, nil
}

func isWordFirst(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isWord(r rune) bool {
	return isWordFirst(r) || unicode.IsDigit(r)
}

// Position ...
func (n *WordNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *WordNode) Equals(m sparse.Node) bool {
	v, ok := m.(*WordNode)
	return ok && v.Value == n.Value
}

// Child ...
func (n *WordNode) Child(i int) sparse.Node {
	panic("Nodes of type 'WordNode' don't have children")
}

// Children ...
func (n *WordNode) Children() int {
	panic("Nodes of type 'WordNode' don't have children")
}

// String ...
func (n *WordNode) String() string {
	return toString("WordNode", n.Row, n.Col, n.Value)
}
