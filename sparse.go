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
		Equals(Node) bool
		Position() (int, int)
	}

	// NodeBuilder ...
	NodeBuilder interface {
		Build() Node
		Add(Node)
		Reset()
	}
)
