package parsers

import (
	"github.com/jlucasnsilva/sparse"
)

type (
	runeParserFunc func(rune) (sparse.Node, error)

	stringParserFunc func(string) (sparse.Node, error)
)

func parseValue(s sparse.Scanner, parse runeParserFunc) (sparse.Scanner, sparse.Node, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}
	r, next := s.Consume()
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	node, err := parse(r)
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
	node, err := parse(value)
	if err != nil {
		return next, nil, err
	}
	return next, node, nil
}
