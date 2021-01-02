package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucasnsilva/sparse"
	"github.com/jlucasnsilva/sparse/parsers"
)

type (
	arrayBuilder struct {
		array []sparse.Node
	}

	array struct {
		array []sparse.Node
	}
)

func main() {
	b1 := newArrayBuilder()
	b2 := newArrayBuilder()
	text := "[1, 2, 3, 4, 5]"
	rdr := strings.NewReader(text)

	s, err := sparse.NewScanner(rdr)
	if err != nil {
		log.Fatalln(err)
	}

	list := sparse.Some(
		parsers.Int,
		parsers.Sequence(", "),
	)
	expr := sparse.And(
		parsers.ThisRune('['),
		list(b1),
		parsers.ThisRune(']'),
	)
	p := expr(b2)

	_, node, err := p(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(node)
}

func newArrayBuilder() *arrayBuilder {
	return &arrayBuilder{array: make([]sparse.Node, 0, 10)}
}

func (a *arrayBuilder) Add(n sparse.Node) {
	a.array = append(a.array, n)
}

func (a *arrayBuilder) Reset() {
	a.array = a.array[:0]
}

func (a *arrayBuilder) Build() sparse.Node {
	return &array{array: a.array}
}

func (a *array) Position() (int, int) {
	return a.array[0].Position()
}

func (a *array) Equals(n sparse.Node) bool {
	// TODO
	return false
}

func (a *array) Child(i int) sparse.Node {
	if i >= 0 && i < len(a.array) {
		return a.array[i]
	}
	return nil
}

func (a *array) Children() int {
	return len(a.array)
}

func (a *array) String() string {
	s := ""
	for _, node := range a.array {
		s += node.String() + "\n"
	}
	return s
}
