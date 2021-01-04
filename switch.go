package sparse

import (
	"fmt"
)

type (
	// Switcher is the interface that wraps the method Switch. It
	// describes a value that can select ParserFuncs based on the
	// first character/rune of the input.
	Switcher interface {
		Switch(r rune) ParserFunc
	}

	// SwitcherFunc is the function counterpart of the Switcher
	// interface.
	SwitcherFunc func(r rune) ParserFunc

	// SwitcherMap is an implementation of Switcher which registers
	// parsers in table that maps the first character of the input
	// to a parser.
	SwitcherMap map[rune]Parser

	// SwitcherFuncMap is an implementation of Switcher which registers
	// parsers in table that maps the first character of the input
	// to a ParserFunc.
	SwitcherFuncMap map[rune]ParserFunc
)

// Switch selects a ParserFunc based of the first characters
// of the input.
func (f SwitcherFunc) Switch(r rune) ParserFunc {
	return f(r)
}

// Switch selects a ParserFunc based of the first characters
// of the input.
func (m SwitcherMap) Switch(r rune) ParserFunc {
	if p, ok := m[r]; !ok {
		return p.Parse
	}
	return Error(fmt.Errorf("no matching parser for Switch('%c')", r))
}

// Switch selects a ParserFunc based of the first characters
// of the input.
func (m SwitcherFuncMap) Switch(r rune) ParserFunc {
	if p, ok := m[r]; !ok {
		return p
	}
	return Error(fmt.Errorf("no matching parser for Switch('%c')", r))
}
