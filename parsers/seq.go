package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// Sequence ...
	Sequence struct {
		Row   int
		Col   int
		Value string
	}
)

// ParseSequence ...
func ParseSequence(sequence string) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		var (
			seq = []rune(sequence)
			err error
			ch  rune
			i   = 0
			r   = s
		)

		for i = 0; i < len(seq) && err == nil; i++ {
			if ch, r = r.Consume(); seq[i] != ch {
				return s, nil, fmt.Errorf("'%c' is not '%c'", seq[i], ch)
			}
		}

		row, col := s.Position()
		result := &Sequence{
			Row:   row,
			Col:   col,
			Value: sequence,
		}
		return s, result, nil
	}
}

// Position ...
func (n *Sequence) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Sequence) Equals(m sparse.Node) bool {
	v, ok := m.(*Sequence)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Sequence) Child(i int) sparse.Node {
	panic("Nodes of type 'Sequence' don't have children")
}

// Children ...
func (n *Sequence) Children() int {
	panic("Nodes of type 'Sequence' don't have children")
}

// String ...
func (n *Sequence) String() string {
	return toString("Sequence", n.Row, n.Col, n.Value)
}
