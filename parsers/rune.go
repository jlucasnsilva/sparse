package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// Rune ...
	Rune struct {
		Row   int
		Col   int
		Value rune
	}
)

// ParseRune ...
func ParseRune(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValue(s, createRune)
}

func createRune(r rune, row, col int) (sparse.Node, error) {
	return &Rune{Value: r, Row: row, Col: col}, nil
}

// ParseThisRune ...
func ParseThisRune(r rune) sparse.ParserFunc {
	pred := func(t rune) bool {
		return t == r
	}
	err := func(t rune) error {
		return fmt.Errorf("expected '%c'. Got '%c' instead", r, t)
	}
	return parseOneRune(pred, err)
}

// ParseOneRune ...
func ParseOneRune(pred func(r rune) bool) sparse.ParserFunc {
	err := func(t rune) error {
		return fmt.Errorf("Invalid character '%c'", t)
	}
	return parseOneRune(pred, err)
}

func parseOneRune(pred func(r rune) bool, err func(rune) error) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		createRune := func(t rune, row, col int) (sparse.Node, error) {
			if pred(t) {
				return &Rune{Value: t, Row: row, Col: col}, nil
			}
			return nil, err(t)
		}
		return parseValue(s, createRune)
	}
}

// Position ...
func (n *Rune) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Rune) Equals(m sparse.Node) bool {
	v, ok := m.(*Rune)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Rune) Child(i int) sparse.Node {
	panic("Nodes of type 'Rune' don't have children")
}

// Children ...
func (n *Rune) Children() int {
	panic("Nodes of type 'Rune' don't have children")
}

// String ...
func (n *Rune) String() string {
	return toString("Rune", n.Row, n.Col, n.Value)
}
