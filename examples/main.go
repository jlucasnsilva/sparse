package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucasnsilva/sparse"
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
		ch := r.Head()
		for r.Err() == nil {
			fmt.Printf("%c", ch)
			ch, r = r.Consume()
		}
		fmt.Println("")
	}

	fmt.Printf("\t%+v\n", s)

	var node sparse.TreeNode
	s, node, err = sparse.Number(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)
	fmt.Printf("\t%+v\n", s)

	s, node, err = sparse.Identifier(s)
	fmt.Printf("node = %+v, error = %v\n", node, err)
	fmt.Printf("\t%+v\n", s)
}
