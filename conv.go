package main

import (
	"flag"
	"fmt"

	"golang.org/x/text/width"
)

func main() {
	var s string
	flag.StringVar(&s, "s", "０１２３４５６７８９", "full-width characters to be converted to half-width characters")
	flag.Parse()
	fmt.Println(width.Narrow.String(s))
}
