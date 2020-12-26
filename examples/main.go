package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jlucasnsilva/sparse"
)

func main() {
	r := strings.NewReader("123.hello,world")
	s, err := sparse.NewScanner(r)
	if err != nil {
		log.Fatalln(err)
	}

	n := sparse.Number{}
	s, err = n.Parse(s)
	if n.IsFloat {
		fmt.Println("float:", n.Float, err)
	} else {
		fmt.Println("int:", n.Int, err)
	}
}
