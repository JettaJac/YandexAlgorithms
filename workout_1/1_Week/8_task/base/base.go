package base

import (
	"fmt"
)

// 1 3 3 2
// 3 1 2 3
// 1 5 1 2
// 5 1 2 1
// 2 2 2 3   5 8
// 0 5 4 4
// 4 4 0 5
// 1 1 4 1
// 2 2 3 3
// 1 2 4 3
// 2 1 3 4
// 1 3 2 2
// 1 2 2 2  4 5

const stop = 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func timeMin(timeDrive, count int) int {
	return (count-1)*(timeDrive+stop) + stop
}
func timeMax(timeDrive, count int) int {
	return (count+1)*(timeDrive+stop) - stop
}

func BasicFunc(a, b, n, m int) {

	minFirst := timeMin(a, n)
	maxFirst := timeMax(a, n)
	minSecond := timeMin(b, m)
	maxSecond := timeMax(b, m)
	// fmt.Println(minFirst, maxFirst)
	// fmt.Println(minSecond, maxSecond)

	oneTime := max(minFirst, minSecond)
	twoTime := min(maxFirst, maxSecond)

	if oneTime <= twoTime {
		fmt.Println(oneTime, twoTime)
	} else {
		fmt.Println(-1)
	}
}

func main() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	BasicFunc(a, b, n, m)
}
