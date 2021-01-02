package parsers

import (
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	boolTestCase struct {
		in     string
		expect bool
		err    bool
	}
)

var (
	boolTestCases = []boolTestCase{
		{"true", true, false},
		{"false", false, false},
		{"FALSE", false, true},
		{"TRUE", false, true},
		{"true1", false, true},
		{"false1", false, true},
	}
)

func TestBool(t *testing.T) {
	for _, test := range boolTestCases {
		test := test
		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := Bool(s)
			b, ok := node.(*BoolNode)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !test.err && !ok:
				t.Errorf("expected a BoolNode, got '%T'", node)
			case !test.err && b.Value != test.expect:
				t.Errorf("expected '%v', got '%v'", test.expect, b.Value)
			}
		})
	}
}
