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
	parse := func(r rune) (sparse.Node, error) {
		return &Rune{Value: r}, nil
	}
	return parseValue(s, parse)
}

// ParseThisRune ...
func ParseThisRune(r rune) sparse.ParserFunc {
	pred := func(t rune) bool {
		return t == r
	}
	err := func(t rune) error {
		return fmt.Errorf("expected '%c'. Got '%c' instead", r, t)
	}
	return parseRune(pred, err)
}

// ParseOneRune ...
func ParseOneRune(pred func(r rune) bool) sparse.ParserFunc {
	err := func(t rune) error {
		return fmt.Errorf("Invalid character '%c'", t)
	}
	return parseRune(pred, err)
}

func parseRune(pred func(r rune) bool, err func(rune) error) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		parse := func(t rune) (sparse.Node, error) {
			if pred(t) {
				return &Rune{Value: t}, nil
			}
			return nil, err(t)
		}
		return parseValue(s, parse)
	}
}

// ParseChar ...
func ParseChar(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	var node sparse.Node
	singleQuoteParser := ParseThisRune('\'')
	r, _, err := singleQuoteParser(s)
	if err != nil {
		return r, nil, err
	}
	r, node, err = ParseRune(r)
	if err != nil {
		return r, nil, err
	}
	r, _, err = singleQuoteParser(r)
	if err != nil {
		return r, nil, err
	}
	return r, node, nil
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

// ValueString ...
func (n *Rune) ValueString() string {
	return fmt.Sprintf("'%c'", n.Value)
}
