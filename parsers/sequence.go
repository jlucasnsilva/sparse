package parsers

import (
	"fmt"

	"github.com/jlucasnsilva/sparse"
)

type (
	// SequenceNode ...
	SequenceNode struct {
		Row   int
		Col   int
		Value string
	}
)

// Sequence ...
func Sequence(sequence string) sparse.ParserFunc {
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
		result := &SequenceNode{
			Row:   row,
			Col:   col,
			Value: sequence,
		}
		return r, result, nil
	}
}

// Position ...
func (n *SequenceNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *SequenceNode) Equals(m sparse.Node) bool {
	v, ok := m.(*SequenceNode)
	return ok && v.Value == n.Value
}

// String ...
func (n *SequenceNode) String() string {
	return toString("SequenceNode", n.Row, n.Col, n.Value)
}
