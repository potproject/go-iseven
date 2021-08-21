package main

import (
	"flag"
	"fmt"
)

//go:generate go run gen/gen.go
func main() {
	flag.Parse()
	a := flag.Args()
	t := IsEven(a[0])
	fmt.Printf("%v", t)
}
