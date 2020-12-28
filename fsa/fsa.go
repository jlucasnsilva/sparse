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

	wordAutomata struct {
		count int
	}

	stringAutomata struct {
		bracket rune
		scape   bool
	}
)

// Configuration ...
var (
	Word = func() Automata {
		return &wordAutomata{}
	}

	Number = func() Automata {
		return &numberAutomata{}
	}

	String = func(bracket rune) Automata {
		return &stringAutomata{bracket: bracket}
	}

	Blank = func() Automata {
		return AutomataFunc(func(r rune) bool {
			return unicode.IsSpace(r) && r != '\n'
		})
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
	return unicode.IsDigit(r) || r == '.' && !a.isFloat
}

// IsValid ...
func (a *wordAutomata) IsValid(r rune) bool {
	if a.count == 0 {
		a.count++
		return isWordFirst(r)
	}
	return isWord(r)
}

func (a *stringAutomata) IsValid(r rune) bool {
	res := r != a.bracket || a.scape
	if r == '\\' {
		a.scape = true
	} else {
		a.scape = false
	}
	return res
}

func isWordFirst(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isWord(r rune) bool {
	return isWordFirst(r) || unicode.IsDigit(r)
}
