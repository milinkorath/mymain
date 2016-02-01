package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sep string
	sep1 := `\r\n`
	fmt.Sscanf(`"`+sep1+`"`, "%q", &sep)
	fmt.Println("----", []byte(sep))
	d := fmt.Sprintf("%q", sep1)
	fmt.Println([]byte(sep))
	fmt.Println("---d", []byte(d))
	fmt.Println(strconv.Quote(sep))
	fmt.Println(strconv.Quote(sep1))
}
