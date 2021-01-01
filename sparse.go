package sparse

import "fmt"

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node Node, err error)

	// Parser ...
	Parser interface {
		Parse(s Scanner) (next Scanner, node Node, err error)
	}

	// Node ...
	Node interface {
		fmt.Stringer
		Position() (int, int)
		Equals(Node) bool
		Child(int) Node
		Children() int
	}

	// NodeBuilder ...
	NodeBuilder interface {
		Build() Node
		Add(Node)
		Reset()
	}

	// DismissNodeBuilder ...
	DismissNodeBuilder struct{}

	// Switcher ...
	Switcher interface {
		Switch(r rune) ParserFunc
	}

	// SwitcherFunc ...
	SwitcherFunc func(r rune) ParserFunc

	// SwitcherMap ...
	SwitcherMap map[rune]Parser

	// SwitcherFuncMap ...
	SwitcherFuncMap map[rune]ParserFunc
)

// Switch ...
func (f SwitcherFunc) Switch(r rune) ParserFunc {
	return f(r)
}

// Build ...
func (b DismissNodeBuilder) Build() Node {
	return nil
}

// Add ...
func (b DismissNodeBuilder) Add(Node) {
	// Do nothing
}

// Reset ...
func (b DismissNodeBuilder) Reset() {
	// Do nothing
}

// Error ...
func Error(err error) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		return s, nil, err
	}
}

// Switch ...
func (m SwitcherMap) Switch(r rune) ParserFunc {
	if p, ok := m[r]; !ok {
		return p.Parse
	}
	return Error(fmt.Errorf("no matching parser for Switch('%c')", r))
}

// Switch ...
func (m SwitcherFuncMap) Switch(r rune) ParserFunc {
	if p, ok := m[r]; !ok {
		return p
	}
	return Error(fmt.Errorf("no matching parser for Switch('%c')", r))
}
