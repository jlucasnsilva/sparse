package parsers

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
	// Float ...
	Float struct {
		Row   int
		Col   int
		Value float64
	}

	// Int ...
	Int struct {
		Row   int
		Col   int
		Value uint64
	}

	// NumberParser ...
	NumberParser struct {
		foundDot bool
		isFloat  bool
	}
)

// Parse ...
func (p *NumberParser) Parse(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, p.check, parseNumber)
}

func (p *NumberParser) check(r rune) bool {
	if p.foundDot {
		p.isFloat = true
	}
	if r == '.' {
		p.foundDot = true
	}
	return unicode.IsDigit(r) || r == '.' && !p.isFloat
}

func parseNumber(value string) (sparse.Node, error) {
	isFloat := func(s string) bool {
		return strings.ContainsRune(s, '.')
	}

	if slen := len(value); slen < 1 || isFloat(value) && slen < 2 {
		return nil, errors.New("Not a number")
	}

	var (
		node sparse.Node
		err  error
	)
	if isFloat(value) {
		fnode := &Float{}
		fnode.Value, err = strconv.ParseFloat(value, 64)
		node = fnode
	} else {
		inode := &Int{}
		inode.Value, err = strconv.ParseUint(value, 10, 64)
		node = inode
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Position ...
func (n *Float) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Float) Equals(m sparse.Node) bool {
	v, ok := m.(*Float)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Float) Child(i int) sparse.Node {
	panic("Nodes of type 'Float' don't have children")
}

// Children ...
func (n *Float) Children() int {
	panic("Nodes of type 'Float' don't have children")
}

// String ...
func (n *Float) String() string {
	return toString("Float", n.Row, n.Col, n.Value)
}

// Position ...
func (n *Int) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *Int) Equals(m sparse.Node) bool {
	v, ok := m.(*Int)
	return ok && v.Value == n.Value
}

// Child ...
func (n *Int) Child(i int) sparse.Node {
	panic("Nodes of type 'Int' don't have children")
}

// Children ...
func (n *Int) Children() int {
	panic("Nodes of type 'Int' don't have children")
}

// String ...
func (n *Int) String() string {
	return toString("Int", n.Row, n.Col, n.Value)
}
