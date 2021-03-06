package sparse

import (
	"errors"
	"fmt"
)

// Or returns a ParserFunc that consumes the input and returns
// the node returned by the first parser among parsers that
// didn't fail (returned error).
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

// Dismiss returns a ParserFunc that returns the input after being
// consumed by parser and a nil node.
func Dismiss(parser ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		r, _, err := parser(s)
		if err != nil {
			return s, nil, err
		}
		return r, nil, nil
	}
}

// Maybe returns a ParserFunc that returns that received input if
// parser fails and returns the consumed input and created node by
// parser otherwise.
func Maybe(parser ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		next, node, err := parser(s)
		if err != nil {
			return s, nil, nil
		}
		return next, node, nil
	}
}

// Switch returns a ParserFunc which selects a parsers registered in
// Switcher based on the first character of the input.
func Switch(switcher Switcher) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		parser := switcher.Switch(s.First())
		return parser(s)
	}
}

// And tries to match the input against all the passed parsers and
// feeds the created nodes to a NodeBuilder. If any of the parsers
// fails, it'll also fail.
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
			result := b.Build()
			b.Reset()
			return r, result, nil
		}
	}
}

// Some will try to match one or more values described by the parser
// 'target' separated by the parser 'separator'.
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
					result := b.Build()
					b.Reset()
					return r, result, nil
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
					result := b.Build()
					b.Reset()
					return r, result, err
				}
				b.Add(node)
			}
			result := b.Build()
			b.Reset()
			return r, result, nil
		}
	}
}

// Pad returns a ParserFunc that consume the input as if it was fed
// to the passed parsers in following order: padding, parser, padding.
func Pad(parser ParserFunc, padding ParserFunc) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		var (
			r    Scanner
			node Node
			err  error
		)

		pad := Maybe(padding)
		r, _, _ = pad(s)
		r, node, err = parser(r)
		if err != nil {
			return s, nil, err
		}
		r, _, _ = pad(r)
		return r, node, nil
	}
}
