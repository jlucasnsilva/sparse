package sparse

import "fmt"

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node Node, err error)

	// Parser ...
	Parser interface {
		Parse(s Scanner) (next Scanner, node Node, err error)
	}

	// Node ...
	Node interface {
		fmt.Stringer
		Position() (int, int)
		Equals(Node) bool
		Child(uint) Node
		Children() uint
	}
)

// Error ...
func Error(err error) ParserFunc {
	return func(s Scanner) (Scanner, Node, error) {
		return s, nil, err
	}
}
