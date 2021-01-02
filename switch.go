package sparse

import (
	"fmt"
)

type (
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
