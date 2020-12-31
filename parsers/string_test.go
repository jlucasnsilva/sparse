package parsers

import (
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	stringParserTestCase struct {
		in      string
		bracket rune
		expect  string
		isError bool
	}
)

var (
	stringParserTestCases = []stringParserTestCase{
		{`"Hello, world!"`, '"', "Hello, world!", false},
		{"Hello world", '"', "", true},
		{"`Hello, world!`", '`', "Hello, world!", false},
		{"%Hello, world\\%!%", '%', "Hello, world%!", false},
		{"", '"', "", true},
	}
)

func TestParseString(t *testing.T) {
	for _, testCase := range stringParserTestCases {
		test := testCase
		t.Run(test.in, func(t *testing.T) {
			expect := String{Value: test.expect, Bracket: test.bracket}

			r := strings.NewReader(test.in)
			s, err := sparse.NewScanner(r)
			if err != nil && !test.isError {
				t.Errorf("Failed to create the scanner: %v", err)
			}

			_, n, err := ParseString(test.bracket)(s)
			switch {
			case !test.isError && err != nil:
				t.Errorf("Got an error when none was expected: %v", err)
			case test.isError && err == nil:
				t.Errorf("Expected and error. Got string '%v'", n)
			case err == nil && !n.Equals(&expect):
				t.Errorf("Expected '%v'. Got '%v'", expect.String(), n)
			}
		})
	}
}
