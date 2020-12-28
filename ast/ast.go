package ast

type (
	// Node ...
	Node interface {
		Position() (int, int)
		ValueString() string
		Equals(Node) bool
		Child(int) Node
		Children() int
	}
)
