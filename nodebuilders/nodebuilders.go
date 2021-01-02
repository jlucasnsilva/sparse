package nodebuilders

import "github.com/jlucasnsilva/sparse"

type (

	// Dismiss ...
	Dismiss struct{}
)

// Build ...
func (b Dismiss) Build() sparse.Node {
	return nil
}

// Add ...
func (b Dismiss) Add(sparse.Node) {
	// Do nothing
}

// Reset ...
func (b Dismiss) Reset() {
	// Do nothing
}
