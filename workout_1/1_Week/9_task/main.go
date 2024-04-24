package main

import (
	"fmt"
	"main/base"
)

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Println(base.BasicFunc(a, b, c, d, e))
}
