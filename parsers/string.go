package parsers

import (
	"fmt"

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

	// StringParser ...
	StringParser struct {
		Bracket rune
		scape   bool
	}
)

// Parse ...
func (p *StringParser) Parse(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	if p.Bracket == 0 {
		p.Bracket = '"'
	}

	var str string
	row, col := s.Position()
	parseFirst := ParseThisRune(p.Bracket)
	r, _, err := parseFirst(s)
	if err != nil {
		return r, nil, err
	}

	str, r = r.ConsumeWhile(p.check)
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
		Bracket: p.Bracket,
	}
	return s, result, nil
}

func (p *StringParser) check(r rune) bool {
	res := r != p.Bracket || p.scape
	if r == '\\' {
		p.scape = true
	} else {
		p.scape = false
	}
	return res
}

// ParseSingleQuoteString ...
func ParseSingleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := StringParser{Bracket: '\''}
	return p.Parse(s)
}

// ParseDoubleQuoteString ...
func ParseDoubleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := StringParser{Bracket: '"'}
	return p.Parse(s)
}

// ParseBackTickString ...
func ParseBackTickString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := StringParser{Bracket: '`'}
	return p.Parse(s)
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
	b := fmt.Sprintf("'%v'", n.Bracket)
	return toString("String", n.Row, n.Col, n.Value, "Bracket", b)
}
