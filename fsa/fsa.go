package fsa

import "unicode"

type (
	// Automata ...
	Automata interface {
		IsValid(r rune) bool
	}

	// AutomataFunc ...
	AutomataFunc func(r rune) bool

	numberAutomata struct {
		isFloat  bool
		foundDot bool
	}

	identifierAutomata struct {
		count int
	}
)

// Configuration ...
var (
	Identifier = func() Automata {
		return AutomataFunc(func(r rune) bool {
			return isIdentifierFirst(r) || unicode.IsDigit(r)
		})
	}

	Number = func() Automata {
		return &numberAutomata{}
	}
)

// IsValid ...
func (a AutomataFunc) IsValid(r rune) bool {
	return a(r)
}

// IsValid ...
func (a *numberAutomata) IsValid(r rune) bool {
	if a.foundDot {
		a.isFloat = true
	}
	if r == '.' {
		a.foundDot = true
	}
	return unicode.IsDigit(r) || r == '.' && !a.foundDot
}

// IsValid ...
func (a *identifierAutomata) IsValid(r rune) bool {
	if a.count == 0 {
		a.count++
		return isIdentifierFirst(r)
	}
	return isIdentifier(r)
}

func isIdentifierFirst(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isIdentifier(r rune) bool {
	return isIdentifierFirst(r) || unicode.IsDigit(r)
}
