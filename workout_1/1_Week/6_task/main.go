package main

// 5 7 3 2
//
import (
	"fmt"
	// "math"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func quare(i, j, z int) int {
	return (i + j) * z
}

func varietes(top1, top2 []int) (res1, res2 int) {
	perS := /*int(math.Inf(1))*/ 2000000
	for i, i2 := 0, 1; i < 2; i, i2 = i+1, i2-1 {
		for j, j2 := 0, 1; j < 2; j, j2 = j+1, j2-1 {
			tmpMax := max(top1[i2], top2[j2])
			tmpS := quare(top1[i], top2[j], tmpMax)
			// fmt.Println(top1[i]+top2[j], max(top1[i2], top2[j2]))
			// fmt.Println(tmpS, perS)0
			if tmpS < perS {
				// fmt.Println("fg")
				perS = tmpS
				res1 = top1[i] + top2[j]
				res2 = tmpMax
			}
		}
	}

	return

}

// func compare(slice []int) (int, int) {
// 	// if slice []
// 	return 0, 0
// }

func main() {
	var a, b, c, d int
	// var slice []int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(a, b, c, d)
	// fmt.Println()
	// fmt.Println(a, b, c, d)
	var top1 = []int{a, b}
	top2 := []int{c, d}
	fmt.Println(top1, top2)
	res1, res2 := varietes(top1, top2)

	fmt.Println(res1, " ", res2)
}

// func main() { // не работает  сканирование с стдио
// 	var tmp []rune

// 	in := bufio.NewReader(os.Stdin)
// 	for {
// 		r, _, err := in.ReadRune()
// 		tmp = append(tmp, r)
// 		if err != io.EOF {
// 			break
// 		}
// 	}
// 	fmt.Println(string(tmp))
// }
