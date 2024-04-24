package main

import (
	"fmt"
	"main/base"
)

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	fmt.Println(base.BasicFunc(n, k, m))
}
