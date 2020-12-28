package ast

type (
	// Node ...
	Node interface {
		Position() (int, int)
		Child(int) Node
	}
)
