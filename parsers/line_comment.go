package parsers

import "github.com/jlucasnsilva/sparse"

type (
	// LineCommentNode ...
	LineCommentNode struct {
		Row   int
		Col   int
		Value string
	}
)

// LineComment ...
func LineComment(start string) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		r, _, err := ThisWord(start)(s)
		if err != nil {
			return s, nil, err
		}

		text, r := r.ConsumeWhile(isNotNewline)
		row, col := s.Position()
		result := &LineCommentNode{
			Row:   row,
			Col:   col,
			Value: text,
		}
		return r, result, nil
	}
}

// DismissLineComment ...
func DismissLineComment(start string) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		r, _, err := ThisWord(start)(s)
		if err != nil {
			return s, nil, err
		}

		_, r = r.ConsumeWhile(isNotNewline)
		return r, nil, nil
	}
}

func isNotNewline(r rune) bool {
	return r != '\n'
}

// Position ...
func (n *LineCommentNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *LineCommentNode) Equals(m sparse.Node) bool {
	v, ok := m.(*LineCommentNode)
	return ok && v.Value == n.Value
}

// Child ...
func (n *LineCommentNode) Child(i int) sparse.Node {
	panic("Nodes of type 'LineCommentNode' don't have children")
}

// Children ...
func (n *LineCommentNode) Children() int {
	panic("Nodes of type 'LineCommentNode' don't have children")
}

// String ...
func (n *LineCommentNode) String() string {
	return toString("LineCommentNode", n.Row, n.Col, n.Value)
}
