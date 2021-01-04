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
	hashtableParser struct {
		valueParser sparse.ParserFunc
	}
)

var (
	hashtableStartParser = sparse.Pad(
		parsers.ThisRune('{'),
		parsers.DismissSpace,
	)

	hashtableEndParsers = sparse.Pad(
		parsers.ThisRune('}'),
		parsers.DismissSpace,
	)

	hashtablePairParserBuilder = sparse.And(
		parsers.DoubleQuoteString,
		sparse.Dismiss(
			sparse.Pad(
				parsers.ThisRune(':'),
				parsers.DismissSpace,
			),
		),
		sparse.Or(
			parsers.DoubleQuoteString,
			parsers.Number,
			parsers.Bool,
		),
	)

	hashtableValueParsersBuilder = sparse.Some(
		hashtablePairParserBuilder(&nodebuilders.Array{}),
		sparse.Pad(
			parsers.ThisRune(','),
			parsers.DismissSpace,
		),
	)
)

func main() {
	// text := ` { "PI"    : 3.14, "hello": "world"      , "flag":true, "inner": { "object": true } }`
	text := `
{
	"PI": 3.14,
	"hello": "world",
	"flag":true,
	"inner": {
		"object": true
	}
}
`

	rdr := strings.NewReader(text)
	s, err := sparse.NewScanner(rdr)
	if err != nil {
		log.Fatalln(err)
	}

	p := hashtableParser{}
	_, node, err := p.Parse(s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(node)
}

func (p *hashtableParser) Parse(s sparse.Scanner) (sparse.Scanner, sparse.Node, error) {
	b := sparse.And(
		hashtableStartParser,
		createHashtableValueParser(),
		hashtableEndParsers,
	)
	parse := b(&nodebuilders.Array{})
	return parse(s)
}

func createHashtableValueParser() sparse.ParserFunc {
	p := hashtableParser{}
	pairParser := sparse.And(
		parsers.DoubleQuoteString,
		sparse.Dismiss(
			sparse.Pad(
				parsers.ThisRune(':'),
				parsers.DismissSpace,
			),
		),
		sparse.Or(
			parsers.DoubleQuoteString,
			parsers.Number,
			parsers.Bool,
			p.Parse,
		),
	)
	b := sparse.Some(
		pairParser(&nodebuilders.Array{}),
		sparse.Pad(
			parsers.ThisRune(','),
			parsers.DismissSpace,
		),
	)
	return b(&nodebuilders.Array{})
}
