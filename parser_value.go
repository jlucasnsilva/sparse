package sparse

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jlucasnsilva/sparse/ast"
	"github.com/jlucasnsilva/sparse/fsa"
)

// Number ...
func Number(s Scanner) (Scanner, ast.Node, error) {
	a := fsa.Number()
	return parseValueWithWhile(s, a.IsValid, parseNumber)
}

// Word ...
func Word(s Scanner) (Scanner, ast.Node, error) {
	a := fsa.Word()
	parse := func(value string) (ast.Node, error) {
		if len(value) < 1 {
			return nil, errors.New("Not an Word")
		}
		return &ast.Word{Value: value}, nil
	}
	return parseValueWithWhile(s, a.IsValid, parse)
}

// Rune ...
func Rune(s Scanner) (Scanner, ast.Node, error) {
	parse := func(r rune) (ast.Node, error) {
		return &ast.Rune{Value: r}, nil
	}
	return parseValue(s, parse)
}

// Blank ...
func Blank(s Scanner) (Scanner, ast.Node, error) {
	a := fsa.Blank()
	parse := func(value string) (ast.Node, error) {
		if len(value) < 1 {
			return nil, errors.New("not white space")
		}
		return &ast.Blank{Value: len(value)}, nil
	}
	return parseValueWithWhile(s, a.IsValid, parse)
}

// Newline ...
func Newline(s Scanner) (Scanner, ast.Node, error) {
	parse := func(r rune) (ast.Node, error) {
		if r != '\n' {
			return nil, errors.New("Not a newline")
		}
		return &ast.Newline{}, nil
	}
	return parseValue(s, parse)
}

func parseValueWithWhile(s Scanner, pred func(rune) bool, parse func(string) (ast.Node, error)) (Scanner, ast.Node, error) {
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

func parseValue(s Scanner, parse func(rune) (ast.Node, error)) (Scanner, ast.Node, error) {
	if err := s.Err(); err != nil {
		return s, nil, err
	}
	r, next := s.Consume()
	if err := next.Err(); err != nil {
		return next, nil, err
	}
	node, err := parse(r)
	if err != nil {
		return next, nil, err
	}
	return next, node, nil
}

func parseNumber(value string) (ast.Node, error) {
	isFloat := func(s string) bool {
		return strings.ContainsRune(s, '.')
	}

	if slen := len(value); slen < 1 || isFloat(value) && slen < 2 {
		return nil, errors.New("Not a number")
	}

	var (
		node ast.Node
		err  error
	)
	if isFloat(value) {
		fnode := &ast.Float{}
		fnode.Value, err = strconv.ParseFloat(value, 64)
		node = fnode
	} else {
		inode := &ast.Int{}
		inode.Value, err = strconv.ParseUint(value, 10, 64)
		node = inode
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}
