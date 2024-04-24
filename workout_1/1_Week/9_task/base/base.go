package base

import (
	"fmt"
	"sort"
)

func sortInt(z []int) []int {

	for i := 0; i < len(z)-1; i++ {
		if z[i] > z[i+1] {
			z[i], z[i+1] = z[i+1], z[i]
		}
	}

	fmt.Println(z)
	return z
}

func BasicFunc(a, b, c, d, e int) string { // wbrk
	var res string
	wals := []int{d, e}
	brick := []int{a, b, c}

	sortInt(wals)
	sortInt(brick)
	sort.Ints(wals)
	sort.Ints(brick)

	if wals[0] >= brick[0] && wals[1] >= brick[1] {
		res = "YES"
	} else {
		res = "NO"
	}

	return res
}

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	fmt.Println(BasicFunc(a, b, c, d, e))
}
