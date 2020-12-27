package sparse

import (
	"errors"
	"strconv"
	"unicode"
)

func parseValueWithWhile(s Scanner, pred func(rune) bool, parse func(string) (TreeNode, error)) (Scanner, TreeNode, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}
	value, next := s.ConsumeWhile(pred)
	if err := next.Err(); err != nil {
		return next, nil, err
	}

	node, err := parse(value)
	if err != nil {
		return next, nil, err
	}
	return next, node, nil
}

// Number ...
func Number(s Scanner) (Scanner, TreeNode, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}

	isFloat := false
	foundDot := false
	isNumber := func(r rune) bool {
		if foundDot {
			isFloat = true
		}
		if r == '.' {
			foundDot = true
		}
		return unicode.IsDigit(r) || r == '.' && !isFloat
	}

	numstr, next := s.ConsumeWhile(isNumber)
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	if slen := len(numstr); slen < 1 || isFloat && slen < 2 {
		return next, nil, errors.New("Not a number")
	}

	var (
		err  error
		node TreeNode
	)
	if isFloat {
		fnode := &FloatNode{}
		fnode.Value, err = strconv.ParseFloat(numstr, 64)
		node = fnode
	} else {
		inode := &IntNode{}
		inode.Value, err = strconv.ParseUint(numstr, 10, 64)
		node = inode
	}
	if err != nil {
		return s, nil, err
	}
	return next, node, nil
}

// Identifier ...
func Identifier(s Scanner) (Scanner, TreeNode, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}

	ok := false
	isIdent := func(r rune) bool {
		if !ok {
			return unicode.IsLetter(r) || r == '_'
		}
		return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
	}

	ident, next := s.ConsumeWhile(isIdent)
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	if len(ident) < 1 {
		return next, nil, errors.New("Not a identifier")
	}
	return next, &IdentifierNode{Value: ident}, nil
}
