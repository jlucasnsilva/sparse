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

	floatTestCase struct {
		in     string
		expect float64
		err    bool
	}

	intTestCase struct {
		in     string
		expect uint64
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

	floatTestCases = []floatTestCase{
		{"3.14", 3.14, false},
		{"3014", 3014, false},
		{"0x14", 0, false},
		{"b101", 0, true},
	}

	intTestCases = []intTestCase{
		{"3.14", 3, false},
		{"3014", 3014, false},
		{"0x14", 0, false},
		{"b101", 0, true},
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

func TestFloat(t *testing.T) {
	for _, test := range floatTestCases {
		test := test

		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := Float(s)
			f, ok := node.(*FloatNode)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !ok && !test.err:
				t.Errorf("expected a FloatNode, got %v of type %T", node, node)
			case !test.err && f.Value != test.expect:
				t.Errorf("expected '%v', got '%v'", test.expect, f.Value)
			}
		})
	}
}

func TestInt(t *testing.T) {
	for _, test := range intTestCases {
		test := test

		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := Int(s)
			i, ok := node.(*IntNode)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !ok && !test.err:
				t.Errorf("expected a IntNode, got %v of type %T", node, node)
			case !test.err && i.Value != test.expect:
				t.Errorf("expected '%v', got '%v'", test.expect, i.Value)
			}
		})
	}
}
