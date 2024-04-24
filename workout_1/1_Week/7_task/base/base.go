package base

import (
	"fmt"
)

func BasicFunc(n, k, m int) int { // wbrk
	var countParts int
	remainder := n
	// i := 0
	if n < k || k < m || k == 0 || m == 0 {
		return 0
	}
	for remainder >= k {
		// fmt.Println(i)
		countWorks := remainder / k
		remainder = remainder % k

		countParts += k / m * countWorks
		remainder += k % m * countWorks
		// fmt.Println(remainder)
		// i++
	}
	return countParts
}

// func BasicFunc(n, k, m int) int { //рекурсия
// 	// var countParts int
// 	countWorks := n / k
// 	remainder := n % k

// 	countParts := k / m * countWorks
// 	remainder += k % m * countWorks

// 	if remainder >= k {
// 		countParts += BasicFunc(remainder, k, m)
// 	}

// 	return countParts
// }

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)

	fmt.Println(BasicFunc(n, k, m))
}
