package parsers

import (
	"errors"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
	// Word ...
	Word struct {
		Row   int
		Col   int
		Value string
	}

	// WordParser ...
	WordParser struct {
		count int
	}
)

// Parse ...
func (p *WordParser) Parse(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, p.check, createWord)
}

func (p *WordParser) check(r rune) bool {
	if p.count == 0 {
		p.count++
		return isWordFirst(r)
	}
	return isWord(r)
}

func createWord(value string, row, col int) (sparse.Node, error) {
	if len(value) < 1 {
		return nil, errors.New("not an word")
	}
	return &Word{Value: value, Row: row, Col: col}, nil
}

func isWordFirst(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isWord(r rune) bool {
	return isWordFirst(r) || unicode.IsDigit(r)
}

// Position ...
func (n *Word) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Word) Equals(m sparse.Node) bool {
	v, ok := m.(*Word)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Word) Child(i int) sparse.Node {
	panic("Nodes of type 'Word' don't have children")
}

// Children ...
func (n *Word) Children() int {
	panic("Nodes of type 'Word' don't have children")
}

// String ...
func (n *Word) String() string {
	return toString("Word", n.Row, n.Col, n.Value)
}
