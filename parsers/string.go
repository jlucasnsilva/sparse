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

	stringStateMachine struct {
		bracket rune
		scape   bool
	}
)

// ParseOneString ...
func ParseOneString(bracket rune) sparse.ParserFunc {
	sm := stringStateMachine{bracket: bracket}
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		var str string

		parseFirst := ParseThisRune(bracket)
		r, _, err := parseFirst(s)
		if err != nil {
			return r, nil, err
		}

		str, r = r.ConsumeWhile(sm.Check)
		if err := r.Err(); err != nil {
			return r, nil, err
		}

		r, _, err = parseFirst(r)
		if err != nil {
			return r, nil, err
		}
		return s, &String{Value: str, Bracket: bracket}, nil
	}
}

// ParseSingleQuoteString ...
func ParseSingleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := ParseOneString('\'')
	return p(s)
}

// ParseDoubleQuoteString ...
func ParseDoubleQuoteString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := ParseOneString('"')
	return p(s)
}

// ParseBackTickString ...
func ParseBackTickString(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	p := ParseOneString('`')
	return p(s)
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

// ValueString ...
func (n *String) ValueString() string {
	return fmt.Sprintf("%c%v%c", n.Bracket, n.Value, n.Bracket)
}

// Check ...
func (sm *stringStateMachine) Check(r rune) bool {
	res := r != sm.bracket || sm.scape
	if r == '\\' {
		sm.scape = true
	} else {
		sm.scape = false
	}
	return res
}
