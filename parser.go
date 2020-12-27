package sparse

type (
	// ParserFunc ...
	ParserFunc func(s Scanner) (next Scanner, node TreeNode, err error)

	// TreeNode ...
	TreeNode interface {
	}
)
