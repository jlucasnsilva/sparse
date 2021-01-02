package parsers

import (
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	lineCommentTestCase struct {
		in     string
		start  string
		expect string
		err    bool
	}
)

var (
	lineCommentTestCases = []lineCommentTestCase{
		{"// hello world", "//", " hello world", false},
		{"// hello world", "// ", "hello world", false},
		{"# hello world", "#", " hello world", false},
		{"# hello world", "# ", "hello world", false},
		{"-- hello world", "--", " hello world", false},
		{"-- hello world", "-- ", "hello world", false},
		{"//hello world", "// ", "", true},
		{"#hello world", "# ", "", true},
		{"# hello world", "// ", "", true},
	}
)

func TestLineComment(t *testing.T) {
	for _, test := range lineCommentTestCases {
		test := test
		t.Run(test.in, func(t *testing.T) {
			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil {
				t.Errorf("Failed to create a scanner: %v", err)
			}

			_, node, err := LineComment(test.start)(s)
			b, ok := node.(*LineCommentNode)
			switch {
			case test.err && err == nil:
				t.Errorf("expected an error when none happened")
			case !test.err && err != nil:
				t.Errorf("no error was expected, got: %v", err)
			case !test.err && !ok:
				t.Errorf("expected a LineCommentNode, got '%T'", node)
			case !test.err && b.Value != test.expect:
				t.Errorf("expected '%v', got '%v'", test.expect, b.Value)
			}
		})
	}
}
