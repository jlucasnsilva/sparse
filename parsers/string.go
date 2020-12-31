package parsers

import (
	"fmt"
	"strings"

	"github.com/jlucasnsilva/sparse"
)

type (
	// String ...
	String struct {
		Row     int
		Col     int
		Bracket rune
		Value   string
	}
)

// ParseString ...
func ParseString(bracket rune) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		var str string
		row, col := s.Position()
		parseFirst := ParseThisRune(bracket)
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
		result := &String{
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

// ParseSingleQuoteString ...
func ParseSingleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return ParseString('\'')(s)
}

// ParseDoubleQuoteString ...
func ParseDoubleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return ParseString('"')(s)
}

// ParseBackTickString ...
func ParseBackTickString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return ParseString('`')(s)
}

// Equals ...
func (n *String) Equals(m sparse.Node) bool {
	v, ok := m.(*String)
	return ok && v.Value == n.Value && v.Bracket == n.Bracket
}

// Child ...
func (n *String) Child(i int) sparse.Node {
	panic("Nodes of type 'String' don't have children")
}

// Children ...
func (n *String) Children() int {
	panic("Nodes of type 'String' don't have children")
}

// Position ...
func (n *String) Position() (int, int) {
	return n.Row, n.Col
}

// String ...
func (n *String) String() string {
	b := fmt.Sprintf("'%c'", n.Bracket)
	return toString("String", n.Row, n.Col, n.Value, "Bracket", b)
}
