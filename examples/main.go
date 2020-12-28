package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucasnsilva/sparse"
	"github.com/jlucasnsilva/sparse/ast"
)

// 01234567890123456
// 123.45hello,world
func main() {
	text := "123.45hello,world"
	r := strings.NewReader(text)
	s, err := sparse.NewScanner(r)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(text)
	{
		r := s
		ch, r := r.Consume()
		for r.Err() == nil {
			fmt.Printf("%c", ch)
			ch, r = r.Consume()
		}
		fmt.Println("")
	}

	var node ast.Node
	s, node, err = sparse.Number(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)

	s, node, err = sparse.Identifier(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)

	prune := sparse.ThisRune(',')
	s, node, err = prune(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)

	s, node, err = sparse.Identifier(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)
}
