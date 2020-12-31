package parsers

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"github.com/jlucasnsilva/sparse"
)

type (
	// FloatNode ...
	FloatNode struct {
		Row   int
		Col   int
		Value float64
	}

	// IntNode ...
	IntNode struct {
		Row   int
		Col   int
		Value uint64
	}
)

// Number ...
func Number(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	isFloat := false
	foundDot := false
	check := func(r rune) bool {
		if foundDot {
			isFloat = true
		}
		if r == '.' {
			foundDot = true
		}
		return unicode.IsDigit(r) || r == '.' && !isFloat
	}
	return parseValueWithWhile(s, check, createNumber)
}

func createNumber(value string, row, col int) (sparse.Node, error) {
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
		fnode := &FloatNode{Row: row, Col: col}
		fnode.Value, err = strconv.ParseFloat(value, 64)
		node = fnode
	} else {
		inode := &IntNode{Row: row, Col: col}
		inode.Value, err = strconv.ParseUint(value, 10, 64)
		node = inode
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Position ...
func (n *FloatNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *FloatNode) Equals(m sparse.Node) bool {
	v, ok := m.(*FloatNode)
	return ok && v.Value == n.Value
}

// Child ...
func (n *FloatNode) Child(i int) sparse.Node {
	panic("Nodes of type 'FloatNode' don't have children")
}

// Children ...
func (n *FloatNode) Children() int {
	panic("Nodes of type 'FloatNode' don't have children")
}

// String ...
func (n *FloatNode) String() string {
	return toString("FloatNode", n.Row, n.Col, n.Value)
}

// Position ...
func (n *IntNode) Position() (int, int) {
	return n.Row, n.Col
}

// Equals ...
func (n *IntNode) Equals(m sparse.Node) bool {
	v, ok := m.(*IntNode)
	return ok && v.Value == n.Value
}

// Child ...
func (n *IntNode) Child(i int) sparse.Node {
	panic("Nodes of type 'IntNode' don't have children")
}

// Children ...
func (n *IntNode) Children() int {
	panic("Nodes of type 'IntNode' don't have children")
}

// String ...
func (n *IntNode) String() string {
	return toString("IntNode", n.Row, n.Col, n.Value)
}
