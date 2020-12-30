package sparse

import "fmt"

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node Node, err error)

	// Parser ...
	Parser interface {
		Parse(s Scanner) (next Scanner, node Node, err error)
	}

	// ExprParserFunc ...
	ExprParserFunc func(s Scanner) (next Scanner, nodes []Node, err error)

	// Node ...
	Node interface {
		fmt.Stringer
		Position() (int, int)
		Equals(Node) bool
		Child(int) Node
		Children() int
	}
)
