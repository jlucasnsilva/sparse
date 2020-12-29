package sparse

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node Node, err error)

	// ExprParserFunc ...
	ExprParserFunc func(s Scanner) (next Scanner, nodes []Node, err error)

	// Node ...
	Node interface {
		Position() (int, int)
		ValueString() string
		Equals(Node) bool
		Child(int) Node
		Children() int
	}

	// StateMachine ...
	StateMachine interface {
		Check(r rune) bool
	}

	// StateMachineFunc ...
	StateMachineFunc func(r rune) bool
)

// Check ...
func (f StateMachineFunc) Check(r rune) bool {
	return f(r)
}
