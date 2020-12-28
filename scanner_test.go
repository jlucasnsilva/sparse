package sparse

import (
	"strings"
	"testing"
	"unicode"
)

type (
	scannerTest struct {
		in    string
		steps []scannerTestStep
	}

	scannerTestStep struct {
		runeOut   rune
		stringOut string
		pos       int
		row       int
		col       int
		err       bool
		runeTest  bool
		exec      execFunc
	}

	execFunc func(Scanner) (rune, string, Scanner)
)

var (
	scannerTest1 = scannerTest{
		in: "abc",
		steps: []scannerTestStep{
			runeStep('a', 1, 0, 1),
			runeStep('b', 2, 0, 2),
			runeStep('c', 3, 0, 3),
			errStep(3, 0, 3),
		},
	}

	scannerTest2 = scannerTest{
		in: "ab\n\ncd\ne",
		steps: []scannerTestStep{
			runeStep('a', 1, 0, 1),
			runeStep('b', 2, 0, 2),
			runeStep('\n', 3, 1, 0),
			runeStep('\n', 4, 2, 0),
			runeStep('c', 5, 2, 1),
			runeStep('d', 6, 2, 2),
			runeStep('\n', 7, 3, 0),
			runeStep('e', 8, 3, 1),
			errStep(8, 3, 1),
		},
	}

	helloLen = len("hello")
	worldLen = len("world")

	scannerTest3 = scannerTest{
		in: "hello\nworld",
		steps: []scannerTestStep{
			stringStep("hello", helloLen, 0, helloLen, unicode.IsLetter),
			runeStep('\n', helloLen+1, 1, 0),
			stringStep("world", helloLen+1+worldLen, 1, worldLen, unicode.IsLetter),
			errStep(helloLen+1+worldLen, 1, worldLen),
		},
	}
)

func TestScanner(t *testing.T) {
	cases := []scannerTest{scannerTest1, scannerTest2, scannerTest3}

	for _, test := range cases {
		t.Run(test.in, func(t *testing.T) {
			rdr := strings.NewReader(test.in)
			scn, err := NewScanner(rdr)
			if err != nil {
				t.Error(err)
			}

			for _, step := range test.steps {
				r, s, next := step.exec(scn)
				t.Run("error", testScannerErr(step, next.Err()))
				t.Run("rune", testScannerRune(step, r))
				t.Run("string", testScannerString(step, s))
				t.Run("position", testScannerPos(step, next.pos))
				t.Run("row", testScannerRow(step, next.row))
				t.Run("col", testScannerCol(step, next.col))
				scn = next
			}
		})
	}
}

func execConsume(s Scanner) (rune, string, Scanner) {
	ch, next := s.Consume()
	return ch, "", next
}

func createExecConsumeWhile(pred func(r rune) bool) execFunc {
	return func(s Scanner) (rune, string, Scanner) {
		str, next := s.ConsumeWhile(pred)
		return 0, str, next
	}
}

func runeStep(r rune, pos, row, col int) scannerTestStep {
	return scannerTestStep{
		runeOut:  r,
		pos:      pos,
		row:      row,
		col:      col,
		runeTest: true,
		exec:     execConsume,
	}
}

func stringStep(s string, pos, row, col int, pred func(r rune) bool) scannerTestStep {
	return scannerTestStep{
		stringOut: s,
		pos:       pos,
		row:       row,
		col:       col,
		exec:      createExecConsumeWhile(pred),
	}
}

func errStep(pos, row, col int) scannerTestStep {
	return scannerTestStep{
		err:  true,
		pos:  pos,
		row:  row,
		col:  col,
		exec: execConsume,
	}
}

func testScannerErr(step scannerTestStep, err error) func(t *testing.T) {
	return func(t *testing.T) {
		if step.err && err == nil {
			t.Error("Expected an error, got <nil>")
		}
	}
}

func testScannerRune(step scannerTestStep, r rune) func(t *testing.T) {
	return func(t *testing.T) {
		if step.runeTest && r != step.runeOut {
			t.Errorf("Expected '%c' got '%c'.", step.runeOut, r)
		}
	}
}

func testScannerString(step scannerTestStep, s string) func(t *testing.T) {
	return func(t *testing.T) {
		if !step.runeTest && s != step.stringOut {
			t.Errorf("Expected '%v' got '%v'.", step.stringOut, s)
		}
	}
}

func testScannerPos(step scannerTestStep, pos int) func(t *testing.T) {
	return func(t *testing.T) {
		if pos != step.pos {
			t.Errorf("Expected '%v' got '%v'.", step.pos, pos)
		}
	}
}

func testScannerRow(step scannerTestStep, row int) func(t *testing.T) {
	return func(t *testing.T) {
		if row != step.row {
			t.Errorf("Expected '%v' got '%v'.", step.row, row)
		}
	}
}

func testScannerCol(step scannerTestStep, col int) func(t *testing.T) {
	return func(t *testing.T) {
		if col != step.col {
			t.Errorf("Expected '%v' got '%v'.", step.col, col)
		}
	}
}
