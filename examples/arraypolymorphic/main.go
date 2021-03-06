package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucasnsilva/sparse"
	"github.com/jlucasnsilva/sparse/nodebuilders"
	"github.com/jlucasnsilva/sparse/parsers"
)

type (
	arrayNodeBuilder struct {
		array sparse.Node
		count int
	}
)

var (
	arrayStartParser = sparse.Pad(
		parsers.ThisRune('['),
		parsers.Space,
	)

	arrayEndParser = sparse.Pad(
		parsers.ThisRune(']'),
		parsers.Space,
	)

	arrayValueDelim = sparse.Pad(
		parsers.ThisRune(','),
		parsers.Space,
	)

	listParserBuilder = sparse.Some(
		sparse.Or(
			parsers.DoubleQuoteString,
			parsers.Number,
			parsers.Bool,
			parsers.Char,
		),
		arrayValueDelim,
	)

	arrayParserBuilder = sparse.And(
		arrayStartParser,
		listParserBuilder(&nodebuilders.Array{}),
		arrayEndParser,
	)

	arrayParser = arrayParserBuilder(&arrayNodeBuilder{})
)

func main() {
	text := `    [1, "Hello, world!", 3.14,     '@'    , true,     false  ]`
	rdr := strings.NewReader(text)
	s, err := sparse.NewScanner(rdr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Parsing the string '%s':\n\n", text)
	_, node, err := arrayParser(s)
	if err != nil {
		log.Fatalln(err)
	}
	if arr, ok := node.(*nodebuilders.ArrayNode); ok {
		for _, v := range arr.Value {
			fmt.Println(v)
		}
	}
}

// Build ...
func (b *arrayNodeBuilder) Build() sparse.Node {
	return b.array
}

// Add ...
func (b *arrayNodeBuilder) Add(n sparse.Node) {
	b.count++
	if b.count == 2 {
		b.array = n
	}
}

// Reset ...
func (b *arrayNodeBuilder) Reset() {
	b.count = 0
	b.array = nil
}
