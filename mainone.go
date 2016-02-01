package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := `"asdsd"`
	fmt.Println(a)
	b := "aaa"
	c := fmt.Sprintf("%q", b)
	fmt.Println(c)
	d := strconv.Quote(b)
	fmt.Println(d)
	e := "\n"
	f := `"` + e + `"`

	fmt.Println([]byte(f))
}
