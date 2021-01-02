package parsers

import "github.com/jlucasnsilva/sparse"

// Error ...
func Error(err error) sparse.ParserFunc {
	return func(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
		return s, nil, err
	}
}
