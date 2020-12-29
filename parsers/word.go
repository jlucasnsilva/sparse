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

	wordStateMachine struct {
		count int
	}
)

// ParseWord ...
func ParseWord(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	sm := &wordStateMachine{}
	parse := func(value string) (sparse.Node, error) {
		if len(value) < 1 {
			return nil, errors.New("Not an Word")
		}
		return &Word{Value: value}, nil
	}
	return parseValueWithWhile(s, sm.Check, parse)
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

// ValueString ...
func (n *Word) ValueString() string {
	return n.Value
}

// Check ...
func (sm *wordStateMachine) Check(r rune) bool {
	if sm.count == 0 {
		sm.count++
		return isWordFirst(r)
	}
	return isWord(r)
}

func isWordFirst(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isWord(r rune) bool {
	return isWordFirst(r) || unicode.IsDigit(r)
}
