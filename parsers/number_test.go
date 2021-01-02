package parsers

import (
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	numberTestCase struct {
		in     string
		expect sparse.Node
		err    bool
	}
)

var (
	numberTestCases = []numberTestCase{
		{"3.14", &FloatNode{Value: 3.14}, false},
		{"3014", &IntNode{Value: 3014}, false},
		{"0x14", &IntNode{Value: 0}, false},
		{"b101", nil, true},
	}
)

func TestNumber(t *testing.T) {
	for _, test := range numberTestCases {
		test := test

		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := Number(s)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !test.err && test.expect.Equals(node):
				t.Errorf("expected '%v', got '%v'", test.expect, node)
			}
		})
	}
}
