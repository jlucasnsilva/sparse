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
	isFloat := strings.ContainsRune(value, '.')
	if slen := len(value); slen < 1 || isFloat && slen < 2 {
		return nil, errors.New("Not a number")
	}

	var (
		node sparse.Node
		fv   float64
		iv   uint64
		err  error
	)
	if isFloat {
		fv, err = strconv.ParseFloat(value, 64)
		if err != nil {
			node = &FloatNode{Row: row, Col: col, Value: fv}
		}
	} else {
		iv, err = strconv.ParseUint(value, 10, 64)
		if err != nil {
			node = &IntNode{Row: row, Col: col, Value: iv}
		}
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Float ...
func Float(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
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
	return parseValueWithWhile(s, check, createFloat)
}

func createFloat(value string, row, col int) (sparse.Node, error) {
	if slen := len(value); slen < 1 || strings.ContainsRune(value, '.') && slen < 2 {
		return nil, errors.New("Not a number")
	}

	fvalue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, err
	}
	result := &FloatNode{Row: row, Col: col, Value: fvalue}
	return result, nil
}

// Int ...
func Int(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	return parseValueWithWhile(s, unicode.IsDigit, createInt)
}

func createInt(value string, row, col int) (sparse.Node, error) {
	if slen := len(value); slen < 1 {
		return nil, errors.New("Not a number")
	}

	ivalue, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return nil, err
	}
	result := &IntNode{Row: row, Col: col, Value: ivalue}
	return result, nil
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

// String ...
func (n *IntNode) String() string {
	return toString("IntNode", n.Row, n.Col, n.Value)
}
