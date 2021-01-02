package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// RuneNode ...
	RuneNode struct {
		Row   int
		Col   int
		Value rune
	}
)

// Rune ...
func Rune(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValue(s, createRune)
}

func createRune(r rune, row, col int) (sparse.Node, error) {
	return &RuneNode{Value: r, Row: row, Col: col}, nil
}

// ThisRune ...
func ThisRune(r rune) sparse.ParserFunc {
	pred := func(t rune) bool {
		return t == r
	}
	err := func(t rune) error {
		return fmt.Errorf("expected '%c'. Got '%c' instead", r, t)
	}
	return parseOneRune(pred, err)
}

// OneRune ...
func OneRune(pred func(r rune) bool) sparse.ParserFunc {
	err := func(t rune) error {
		return fmt.Errorf("Invalid character '%c'", t)
	}
	return parseOneRune(pred, err)
}

func parseOneRune(pred func(r rune) bool, err func(rune) error) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		createRune := func(t rune, row, col int) (sparse.Node, error) {
			if pred(t) {
				return &RuneNode{Value: t, Row: row, Col: col}, nil
			}
			return nil, err(t)
		}
		return parseValue(s, createRune)
	}
}

// Position ...
func (n *RuneNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *RuneNode) Equals(m sparse.Node) bool {
	v, ok := m.(*RuneNode)
	return ok && v.Value == n.Value
}

// String ...
func (n *RuneNode) String() string {
	v := fmt.Sprintf("%c", n.Value)
	return toString("RuneNode", n.Row, n.Col, v)
}
