package sparse

type (
	// NodeBuilder ...
	NodeBuilder interface {
		Build() Node
		Add(Node)
		Reset()
	}

	// DismissNodeBuilder ...
	DismissNodeBuilder struct{}
)

// Build ...
func (b DismissNodeBuilder) Build() Node {
	return nil
}

// Add ...
func (b DismissNodeBuilder) Add(Node) {
	// Do nothing
}

// Reset ...
func (b DismissNodeBuilder) Reset() {
	// Do nothing
}
