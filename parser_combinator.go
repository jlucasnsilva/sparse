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

// Switch ...
func Switch(switcher Switcher) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		parser := switcher.Switch(s.First())
		return parser(s)
	}
}

// And ...
func And(parsers ...ParserFunc) func(NodeBuilder) ParserFunc {
	return func(b NodeBuilder) ParserFunc {
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
}

// Some ...
func Some(target ParserFunc, separator ParserFunc) func(NodeBuilder) ParserFunc {
	return func(b NodeBuilder) ParserFunc {
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
			t = r

			for {
				if t, node, err = separator(t); err != nil {
					return r, b.Build(), nil
				}
				if t, node, err = target(t); err != nil {
					return r, nil, errors.New("expression ended on a separator")
				}
				b.Add(node)
				r = t
			}
		}
	}

}

// Concat ...
func Concat(parsers ...ParserFunc) func(NodeBuilder) ParserFunc {
	return func(b NodeBuilder) ParserFunc {
		return func(s Scanner) (Scanner, Node, error) {
			var (
				node Node
				err  error
				r    = s
			)
			for _, p := range parsers {
				r, node, err = p(r)
				if err != nil {
					return r, b.Build(), err
				}
				b.Add(node)
			}
			return r, b.Build(), nil
		}
	}
}
