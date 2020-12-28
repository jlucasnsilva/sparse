package sparse

import (
	"fmt"

	"github.com/jlucasnsilva/sparse/ast"
	"github.com/jlucasnsilva/sparse/fsa"
)

// ThisRune ...
func ThisRune(r rune) ParserFunc {
	pred := func(t rune) bool {
		return t == r
	}
	err := func(t rune) error {
		return fmt.Errorf("expected '%c'. Got '%c' instead", r, t)
	}
	return parseRune(pred, err)
}

// OneRune ...
func OneRune(pred func(r rune) bool) ParserFunc {
	err := func(t rune) error {
		return fmt.Errorf("Invalid character '%c'", t)
	}
	return parseRune(pred, err)
}

// OneString ...
func OneString(bracket rune) ParserFunc {
	a := fsa.String(bracket)
	return func(s Scanner) (Scanner, ast.Node, error) {
		var str string

		parseFirst := ThisRune(bracket)
		r, _, err := parseFirst(s)
		if err != nil {
			return r, nil, err
		}

		str, r = r.ConsumeWhile(a.IsValid)
		if err := r.Err(); err != nil {
			return r, nil, err
		}

		r, _, err = parseFirst(r)
		if err != nil {
			return r, nil, err
		}
		return s, &ast.String{Value: str, Bracket: bracket}, nil
	}
}

func parseRune(pred func(r rune) bool, err func(rune) error) ParserFunc {
	return func(s Scanner) (Scanner, ast.Node, error) {
		parse := func(t rune) (ast.Node, error) {
			if pred(t) {
				return &ast.Rune{Value: t}, nil
			}
			return nil, err(t)
		}
		return parseValue(s, parse)
	}
}
