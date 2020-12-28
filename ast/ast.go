package ast

type (
	// Node ...
	Node interface {
		Position() (int, int)
		Equals(Node) bool
		Child(int) Node
		Children() int
	}
)
