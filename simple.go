package sparse

import (
	"errors"
	"strconv"
	"unicode"
)

type (
	// Number ...
	Number struct {
		Int      uint64
		Float    float64
		IsFloat  bool
		foundDot bool
	}
)

// Parse ...
func (p *Number) Parse(s Scanner) (Scanner, error) {
	numstr, next := s.ConsumeWhile(p.isNumber)
	if slen := len(numstr); slen < 1 || p.foundDot && slen < 2 {
		return s, errors.New("Not a number")
	}

	var err error
	if p.foundDot {
		p.Float, err = strconv.ParseFloat(numstr, 64)
		p.IsFloat = true
	} else {
		p.Int, err = strconv.ParseUint(numstr, 10, 64)
	}
	if err != nil {
		return s, err
	}
	return next, nil
}

func (p *Number) isNumber(r rune) bool {
	if r == '.' {
		p.foundDot = true
	}
	return unicode.IsDigit(r) || r == '.' && !p.foundDot
}
