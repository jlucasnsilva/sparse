package parsers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/jlucasnsilva/sparse"
)

type (
	parseValueTest struct {
		in    string
		units []parseValueTestUnit
	}

	parseValueTestUnit struct {
		label  string
		parser ParserFunc
		check  func(sparse.Node) (bool, string, string)
	}
)

var (
	parseValueTest1 = parseValueTest{
		in: ">    let number = 12.34",
		units: []parseValueTestUnit{
			{"Rune", Rune, createTestRune('>')},
			{"Blank", Blank, createTestBlank(4)},
			{"Word", Word, createTestWord("let")},
			{"Blank", Blank, createTestBlank(1)},
			{"Word", Word, createTestWord("number")},
			{"Blank", Blank, createTestBlank(1)},
			{"Rune", Rune, createTestRune('=')},
			{"Blank", Blank, createTestBlank(1)},
			{"Float", Number, createTestFloat(12.34)},
		},
	}
)

func TestParseValue(t *testing.T) {
	cases := []parseValueTest{parseValueTest1}

	for _, test := range cases {
		rdr := strings.NewReader(test.in)
		scn, err := NewScanner(rdr)
		if err != nil {
			t.Error(err)
		}

		t.Run(test.in, func(t *testing.T) {
			for _, u := range test.units {
				next, node, err := u.parser(scn)
				t.Run(u.label, checkUnit(u.check, node, err))
				scn = next
			}
		})
	}
}

func checkUnit(check func(sparse.Node) (bool, string, string), node sparse.Node, err error) func(*testing.T) {
	return func(t *testing.T) {
		if err != nil {
			t.Error(err)
		} else if ok, expected, got := check(node); !ok {
			t.Errorf("Expected '%v' got '%v'\n", expected, got)
		}
	}
}

func createTestFloat(expect float64) func(sparse.Node) (bool, string, string) {
	return func(n sparse.Node) (bool, string, string) {
		_, ok := n.(*Float)
		e := Float{Value: expect}
		return ok && n.Equals(&e), fmt.Sprint(expect), n.ValueString()
	}
}

func createTestUint(expect uint64) func(sparse.Node) (bool, string, string) {
	return func(n sparse.Node) (bool, string, string) {
		_, ok := n.(*Int)
		e := Int{Value: expect}
		return ok && n.Equals(&e), fmt.Sprint(expect), n.ValueString()
	}
}

func createTestWord(expect string) func(sparse.Node) (bool, string, string) {
	return func(n sparse.Node) (bool, string, string) {
		_, ok := n.(*Word)
		e := Word{Value: expect}
		return ok && n.Equals(&e), fmt.Sprint(expect), n.ValueString()
	}
}

func createTestRune(expect rune) func(sparse.Node) (bool, string, string) {
	return func(n sparse.Node) (bool, string, string) {
		_, ok := n.(*Rune)
		e := Rune{Value: expect}
		return ok && n.Equals(&e), fmt.Sprint(expect), n.ValueString()
	}
}

func createTestBlank(expect int) func(sparse.Node) (bool, string, string) {
	return func(n sparse.Node) (bool, string, string) {
		_, ok := n.(*Blank)
		e := Blank{Value: expect}
		return ok && n.Equals(&e), fmt.Sprintf("[blank:%v]", expect), n.ValueString()
	}
}
