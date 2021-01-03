package sparse

import (
	"bytes"
	"errors"
	"fmt"
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

// First ...
func (s Scanner) First() rune {
	return s.text[s.pos]
}

// Consume ...
func (s Scanner) Consume() (ch rune, next Scanner) {
	if s.pos == len(s.text) {
		s.err = fmt.Errorf("end of file (L%v C%v)", s.row, s.col)
	}
	if s.err != nil {
		return 0, s
	}

	ch = s.First()

	s.pos++
	if ch == '\n' {
		s.col = 0
		s.row++
	} else {
		s.col++
	}
	return ch, s
}

// ConsumeWhile ...
func (s Scanner) ConsumeWhile(pred func(rune) bool) (string, Scanner) {
	if s.err != nil {
		return "", s
	}

	tlen := len(s.text)
	if s.pos == tlen {
		s.err = fmt.Errorf("end of file (L%v C%v)", s.row, s.col)
		return "", s
	}

	ch := s.First()
	start := s.pos
	for pred(ch) && s.pos < tlen {
		if ch == '\n' {
			s.col = 0
			s.row++
		} else {
			s.col++
		}

		s.pos++
		if s.pos < tlen {
			ch = s.First()
		}
	}
	return string(s.text[start:s.pos]), s
}

// String ...
func (s Scanner) String() string {
	return fmt.Sprintf(
		"text: %v, pos: %v (L%v, C%v), err = %v",
		string(s.text), s.pos, s.row, s.col, s.err,
	)
}
