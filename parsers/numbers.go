package parsers

import (
	"errors"
	"fmt"
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

	numberStateMachine struct {
		foundDot bool
		isFloat  bool
	}
)

// ParseNumber ...
func ParseNumber(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	sm := &numberStateMachine{}
	return parseValueWithWhile(s, sm.Check, parseNumber)
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

// ValueString ...
func (n *Float) ValueString() string {
	return fmt.Sprint(n.Value)
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

// ValueString ...
func (n *Int) ValueString() string {
	return fmt.Sprint(n.Value)
}

// Check ...
func (sm *numberStateMachine) Check(r rune) bool {
	if sm.foundDot {
		sm.isFloat = true
	}
	if r == '.' {
		sm.foundDot = true
	}
	return unicode.IsDigit(r) || r == '.' && !sm.isFloat
}
