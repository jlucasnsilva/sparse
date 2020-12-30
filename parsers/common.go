package parsers

import (
	"fmt"
	"strings"

	"github.com/jlucasnsilva/sparse"
)

type (
	runeParserFunc func(rune, int, int) (sparse.Node, error)

	stringParserFunc func(string, int, int) (sparse.Node, error)
)

func parseValue(s sparse.Scanner, parse runeParserFunc) (sparse.Scanner, sparse.Node, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}
	r, next := s.Consume()
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	row, col := s.Position()
	node, err := parse(r, row, col)
	if err != nil {
		return next, nil, err
	}
	return next, node, nil
}

func parseValueWithWhile(s sparse.Scanner, pred func(rune) bool, parse stringParserFunc) (sparse.Scanner, sparse.Node, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}
	value, next := s.ConsumeWhile(pred)
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	row, col := s.Position()
	node, err := parse(value, row, col)
	if err != nil {
		return next, nil, err
	}
	return next, node, nil
}

func toString(nodeType string, row, col int, value interface{}, pairs ...interface{}) string {
	b := strings.Builder{}
	for i := 0; i < len(pairs)-1; i += 2 {
		fmt.Fprintf(&b, ", %v: %v", pairs[i], pairs[i+1])
	}
	return fmt.Sprintf(
		"%v{ Row: %v, Col: %v, Value «%v»%v }",
		nodeType, row, col, value, b.String(),
	)
}
