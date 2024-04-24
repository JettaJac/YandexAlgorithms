package main

import (
	"fmt"
	"math"
)

// 1 0 0  0
// 1 0 -1 NS
// 0 2 3 NS
// 0 2 2 MS
// 3 -2 1  1
// 3 3 1 NS
// -5 1 1
// -1 -5 1  -6
// 3 -2 3 NS

func checkInt(number float32) bool {
	// result := int((number * 10)) % 10
	// // fmt.Println(number, "  ", result, "  ", number*10)
	// if result == 0 {
	// 	return true
	// }
	// return false
	return math.Mod(float64(number), 1) == 0
}

func equation(a, b, c int, result *int) int {
	answer := 0
	var tmpResult float32
	// if a == 0 {
	// 	if c * c - b == 0
	if c < 0 || a == 0 && b != c {
		answer = 1
	} else if a == 0 && b == c*c {
		answer = 2
	} else {
		tmpResult = (float32((c * c) - b)) / float32(a)
		// fmt.Println("Tmp ", tmpResult)
		if !checkInt(tmpResult) {
			answer = 1
		} else {
			*result = int(tmpResult)
			// fmt.Println("result  ", *result)
		}
		// fmt.Println(!checkInt(tmpResult))
		// fmt.Println("Test ", checkInt(12.12), " ", checkInt(12))
		// fmt.Println(*result)
	}

	if (a**result + b) < 0 {
		answer = 1
	}

	return answer
}

func main() {
	var a, b, c, result int
	fmt.Scan(&a, &b, &c)
	tmp := equation(a, b, c, &result)
	if tmp == 1 {
		fmt.Println("NO SOLUTION")
	} else if tmp == 2 {
		fmt.Println("MANY SOLUTIONS")
	} else if tmp == 0 {
		fmt.Println(result)
	}
}

// func main() {
// 	var a, b, c int
// 	fmt.Scan(&a, &b, &c)

// 	if c < 0 || (a == 0 && b != c) {
// 		fmt.Println("NO SOLUTION")
// 		return
// 	}

// 	if a == 0 && b == c {
// 		fmt.Println("MANY SOLUTIONS")
// 		return
// 	}

// 	if a == 0 {
// 		fmt.Println("NO SOLUTION")
// 		return
// 	}

// 	discriminant1 := /*math.Sqrt(*/ (float32(c*c - b)) / float32(a)
// 	// fmt.Println("D1  ", discriminant1)
// 	if !checkInt(discriminant1) {
// 		fmt.Println("NO SOLUTION")
// 		return
// 	}
// 	discriminant := int(discriminant1)
// 	// fmt.Println("D0  ", discriminant)
// 	// fmt.Println(math.Sqrt(float64(a*discriminant + b)))
// 	// fmt.Println(discriminant, (a*discriminant+b) == c*c, (a*discriminant+b) >= 0)
// 	if ((a*discriminant + b) == c*c) && (a*discriminant+b) >= 0 {
// 		fmt.Println(discriminant)
// 	} else {
// 		fmt.Println("NO SOLUTION")
// 	}
// }
