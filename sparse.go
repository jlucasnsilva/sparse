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
		Child(int) Node
		Children() int
	}

	// NodeBuilder ...
	NodeBuilder interface {
		Build() Node
		Add(Node)
		Reset()
	}
)
