package sparse

import (
	"bytes"
	"errors"
	"io"
)

type (
	// Scanner ...
	Scanner struct {
		text []rune
		err  error
		pos  int
		row  int
		col  int
	}
)

const delimiters = ",.;[](){}\"'`"

// NewScanner ...
func NewScanner(rdr io.Reader) (Scanner, error) {
	scn := Scanner{}
	b := bytes.Buffer{}
	_, err := b.ReadFrom(rdr)
	if err != nil {
		return scn, err
	}
	s := b.String()
	if len(s) < 1 {
		return scn, errors.New("Empty source")
	}

	scn.text = []rune(s)
	return scn, nil
}

// Position ...
func (s Scanner) Position() (row int, col int) {
	return s.row, s.col
}

// Err ...
func (s Scanner) Err() error {
	return s.err
}

// Head ...
func (s Scanner) Head() rune {
	return s.text[s.pos]
}

// Consume ...
func (s Scanner) Consume() (ch rune, next Scanner) {
	if s.err != nil {
		return 0, s
	}

	n := s
	n.pos++
	if n.pos == len(n.text) {
		n.err = io.EOF
		return 0, n
	}

	ch = n.Head()
	if ch == '\n' {
		n.col = 0
		n.row++
	} else {
		n.col++
	}
	return ch, n
}

// ConsumeWhile ...
func (s Scanner) ConsumeWhile(pred func(rune) bool) (string, Scanner) {
	if s.err != nil {
		return "", s
	}

	tlen := len(s.text)
	n := s
	n.pos++
	if n.pos == tlen {
		n.err = io.EOF
		return "", n
	}

	ch := n.Head()
	start := n.col
	for pred(ch) && n.pos < tlen {
		if ch == '\n' {
			n.col = 0
			n.row++
		} else {
			n.col++
		}

		n.pos++
		if n.pos < tlen {
			ch = n.Head()
		}
	}

	if n.pos == tlen {
		n.pos--
	}
	return string(n.text[start : n.col+1]), n
}
