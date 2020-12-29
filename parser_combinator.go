package sparse

import (
	"errors"
	"fmt"
)

// And ...
func And(parsers ...ParserFunc) ExprParserFunc {
	return func(s Scanner) (Scanner, []Node, error) {
		var (
			nodes = make([]Node, 0, 10)
			node  Node
			err   error
			r     = s
		)
		for i, p := range parsers {
			r, node, err = p(r)
			if err != nil {
				return s, nil, fmt.Errorf("invalid parser at %v: %v", i, err)
			}
			nodes = append(nodes, node)
		}
		return r, nodes, nil
	}
}

// Or ...
func Or(parsers ...ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		for _, p := range parsers {
			if r, node, err := p(s); err == nil {
				return r, node, nil
			}
		}
		return s, nil, errors.New("No matching parser")
	}
}

// Some ...
func Some(target ParserFunc, separator ParserFunc) ExprParserFunc {
	return func(s Scanner) (Scanner, []Node, error) {
		var (
			nodes = make([]Node, 0, 10)
			node  Node
			err   error
			r     = s
			t     = r
		)

		if r, node, err = target(r); err != nil {
			return s, nil, errors.New("not a single match")
		}
		nodes = append(nodes, node)

		for true {
			if t, node, err = target(r); err != nil {
				return r, nodes, nil
			}
			if t, node, err = target(t); err != nil {
				return r, nil, errors.New("expression ended on a separator")
			}
			nodes = append(nodes, node)
			r = t
		}
		return s, nil, nil
	}
}
