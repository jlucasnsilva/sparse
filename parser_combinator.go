package sparse

import (
	"errors"
	"fmt"
)

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

// And ...
func And(b NodeBuilder, parsers ...ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		var (
			node Node
			err  error
			r    = s
		)
		for i, p := range parsers {
			r, node, err = p(r)
			if err != nil {
				b.Reset()
				return s, nil, fmt.Errorf("invalid parser at %v: %v", i, err)
			}
			b.Add(node)
		}
		return r, b.Build(), nil
	}
}

// Some ...
func Some(b NodeBuilder, target ParserFunc, separator ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		var (
			node Node
			err  error
			r    = s
			t    = r
		)

		if r, node, err = target(r); err != nil {
			return s, nil, errors.New("not a single match")
		}
		b.Add(node)

		for true {
			if t, node, err = target(r); err != nil {
				return r, b.Build(), nil
			}
			if t, node, err = target(t); err != nil {
				return r, nil, errors.New("expression ended on a separator")
			}
			b.Add(node)
			r = t
		}
		return s, nil, nil
	}
}
