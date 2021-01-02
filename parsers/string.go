package parsers

import (
	"fmt"
	"strings"

	"github.com/jlucasnsilva/sparse"
)

type (
	// StringNode ...
	StringNode struct {
		Row     int
		Col     int
		Bracket rune
		Value   string
	}
)

// String ...
func String(bracket rune) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		var str string
		row, col := s.Position()
		parseFirst := ThisRune(bracket)
		r, _, err := parseFirst(s)
		if err != nil {
			return r, nil, err
		}

		str, r = consumeString(r, bracket)
		if err := r.Err(); err != nil {
			return r, nil, err
		}

		r, _, err = parseFirst(r)
		if err != nil {
			return r, nil, err
		}
		result := &StringNode{
			Col:     col,
			Row:     row,
			Value:   str,
			Bracket: bracket,
		}
		return s, result, nil
	}
}

func consumeString(s sparse.Scanner, bracket rune) (string, sparse.Scanner) {
	if s.Err() != nil {
		return "", s
	}

	var (
		ch rune
		r  sparse.Scanner
	)
	scape := false
	ch, s = s.Consume()
	b := strings.Builder{}
	for r.Err() == nil && (ch != bracket || scape) {
		r = s
		if ch == '\\' {
			scape = true
		} else {
			scape = false
			b.WriteRune(ch)
		}
		ch, s = s.Consume()
	}
	return b.String(), r
}

// SingleQuoteString ...
func SingleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return String('\'')(s)
}

// DoubleQuoteString ...
func DoubleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return String('"')(s)
}

// BackTickString ...
func BackTickString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return String('`')(s)
}

// Equals ...
func (n *StringNode) Equals(m sparse.Node) bool {
	v, ok := m.(*StringNode)
	return ok && v.Value == n.Value && v.Bracket == n.Bracket
}

// Child ...
func (n *StringNode) Child(i uint) sparse.Node {
	panic("Nodes of type 'StringNode' don't have children")
}

// Children ...
func (n *StringNode) Children() uint {
	panic("Nodes of type 'StringNode' don't have children")
}

// Position ...
func (n *StringNode) Position() (int, int) {
	return n.Row, n.Col
}

// StringNode ...
func (n *StringNode) String() string {
	b := fmt.Sprintf("'%c'", n.Bracket)
	return toString("StringNode", n.Row, n.Col, n.Value, "Bracket", b)
}
