package parsers

import (
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	charTestCase struct {
		in     string
		expect rune
		err    bool
	}
)

var (
	charTestCases = []charTestCase{
		{"'a'", 'a', false},
		{"'&'", '&', false},
		{`"a"`, 0, true},
		{"`a`", 0, true},
		{"a", 0, true},
	}
)

func TestChar(t *testing.T) {
	for _, test := range charTestCases {
		test := test
		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := Char(s)
			b, ok := node.(*CharNode)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !test.err && !ok:
				t.Errorf("expected a CharNode, got '%T'", node)
			case !test.err && b.Value != test.expect:
				t.Errorf("expected '%v', got '%v'", test.expect, b.Value)
			}
		})
	}
}
