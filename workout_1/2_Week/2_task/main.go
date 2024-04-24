package main

import (
	"bufio"
	"fmt"
	"os"
	// "os"
)

//	func myScanInt(n int, input *os.File) (res []int) {
//		var c int
//		// fmt.Println(n, c)
//		for i := 0; i < n; i++ {
//			fmt.Scan(&c)
//			// fmt.Println(c)
//			res = append(res, c)
//		}
//		return
//	}
func myScanString(n int, input *os.File) (res []string) {
	var c string
	// fmt.Println(n, c)
	for i := 0; i < n; i++ {
		fmt.Scan(&c)
		// fmt.Println(c)
		res = append(res, c)
	}
	return
}

// func main() {
// 	var a, b, c, d, res int
// 	// if checkFuncS(readStdinInt(os.Stdin)) {
// 	// 	fmt.Println("YES")
// 	// 	return
// 	// }
// 	// fmt.Println("NO")
// 	fmt.Scan(&a, &b, &c, &d)
// 	tmp := d - b
// 	// fmt.Println(tmp)
// 	if tmp < 0 {
// 		tmp = 0
// 	}
// 	res = a + tmp*c

//		fmt.Println(res)
//	}

/*
	func changeFunc(tmp []int) (count int) {
		var min0, min2 int
		for i := 0; i < len(tmp)-1; i++ {
			if tmp[i]%10 == 0 {
				min0 = i
				tmp[i] = 0
			} else if tmp[min0] > tmp[i+1] && tmp[i+1] < 10 {
				min0 = i + 1

			} else if tmp[min0] > tmp[i+1] && tmp[i+1] > 10 {
				// min = tmp[i+1]
				min2 = i + 1
				fmt.Println(" ttt", count, min2)
			}

		}
		return
	}

	func main() {
		var a, b, c, res int

		fmt.Scan(&a, &b)

		tmp := make([]int, a)

		for i := 0; i < a; i++ {
			fmt.Scan(&c)
			tmp[i] = c
		}
		fmt.Println("rrr ", tmp)
		for i := 0; i < b; i++ {
			count := changeFunc(tmp)
			// // j := count
			// fmt.Println("old ", tmp[1])
			res += (9 - tmp[count])
			// tmp[count] = 9
			// // 	// tmp(j) = 9
			fmt.Println(" chang ", tmp[count])
		}

		fmt.Println(res)
	}
*/
/*func main() {
	var sum, a, b, c, res int
	fmt.Scan(&sum, &a, &b, &c)
	m := 1
	for {
		if tmp < sum {
			res++
		}
	}

	fmt.Prinln(res)
}*/
/* обеды и купоны
func changeSlice(slice []int) {
	max := 0
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			max = i + 1
		}
	}
	slice[max] = 0
}

func baseFunc(slice []int) (res int) {
	var kupon uint
	for i := 0; i < len(slice); i++ {
		if kupon > 0 {
			// tmp := slice[i:]
			changeSlice(slice[i:])
			kupon--
		}
		if slice[i] > 100 {
			kupon++
		}
		res += slice[i]
	}
	return
}*/

// func baseFunc(slice []int) (res []int) { положение четной на не четной
// 	var count, count2, x int
// 	res = append(res, 0, 0, 0)
// 	for i := 0; i < len(slice); i++ {
// 		if slice[i]%2 == 0 && (i+1)%2 == 1 {
// 			res[x] = i + 1
// 			count++
// 			count2++
// 			x++
// 		}
// 		if slice[i]%2 == 1 && (i+1)%2 == 0 {
// 			res[x] = i + 1
// 			count--
// 			count2++
// 			x++
// 		}

// 		if count > 1 || count < -1 || count2 > 2 || count2 == 0 {
// 			res[0] = -1
// 			res[1] = -1
// 			break
// 		}
// 	}
// 	return
// }
/*
func calcFunc(slice []int) (res int) {
	for i := 0; i < 7; i++ {
		if slice[i] == 5 {
			res++
		}
		if slice[i] == 2 || slice[i] == 3 {
			res = -1
			break
		}
	}
	return
}

func baseFunc(slice []int) (res int) {
	storage := []int{}
	for i := 0; i < len(slice)-6; i++ {
		newSlice := slice[i:]
		storage = append(storage, calcFunc(newSlice))

	}
	// fmt.Println(storage)
	sort.Ints(storage)
	res = storage[len(storage)-1]

	return
}

func main() {
	var n  int
	fmt.Scan(&n)
	mySlice := myScan(n, os.Stdin)
	// fmt.Println(mySlice)

	res := baseFunc(mySlice)
	// fmt.Println(res[0], res[1])
	fmt.Println(res)
}
*/

// func calcFunc(slice []int) (res int) {
// 	for i := 0; i < 7; i++ {
// 		if slice[i] == 5 {
// 			res++
// 		}
// 		if slice[i] == 2 || slice[i] == 3 {
// 			res = -1
// 			break
// 		}
// 	}
// 	return
// }

// func sortMySlice(storage [][]string) {
// 	sort.Slice(storage, func(i, j int) bool {
// 		// Сравниваем значения первых элементов в срезах
// 		return storage[i][0] < storage[j][0]
// 	})

// 	sort.Slice(storage, func(i, j int) bool {
// 		// Сравниваем значения первых элементов в срезах
// 		return storage[i][0] == storage[j][0]
// 	})
// 	fmt.Println(storage)

// }

// 4
// b/b/c/d
// a/b
// a/b/c
// a
/*
func myPrint(line string) {
	fmt.Println("  ", line)
}

func baseFunc(slice []string) (res string) {
	var storage [][]string

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		newSlice := strings.Split(slice[i], "/")
		// fmt.Println(newSlice)
		storage = append(storage, newSlice)

	}
	// fmt.Println()
	// sortMySlice(storage)
	// fmt.Println(storage)
	printPath(storage)

	// sort.slice(storage)
	// res = storage[len(storage)-1]

	return
}
func printTmp(slice string) {
	fmt.Println("  ", slice)

}

func printPath(slice [][]string) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if j == 0 {
				fmt.Println(slice[i][j])
			} else {
				for x := 0; x < j; x++ {
					fmt.Print("  ")
				}
				fmt.Println(slice[i][j])
			}

		}
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	mySlice := myScanString(n, os.Stdin)
	// fmt.Println(mySlice)

	baseFunc(mySlice)
	// fmt.Println(res[0], res[1])
	// fmt.Println(res)
}
*/
/*

func changeMatrix(matrix [][]int) [][]int {
	x := len(matrix)
	y := len(matrix[0])

	changed := make([][]int, y)
	for i := range changed {
		changed[i] = make([]int, x)
	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			changed[j][x-i-1] = matrix[i][j]
		}
	}

	return changed
}

func matrixPrint(x, y int, matrix [][]int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	matrix := make([][]int, x)
	for i := 0; i < x; i++ {
		matrix[i] = make([]int, y)
		for j := 0; j < y; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}
	matrixPrint(x, y, changeMatrix(matrix))
}
*/
/*
func pathForest(forest [][]string, i, j int, visited [][]bool, test *string) int {

	if !chackPath(forest, i, j) || visited[i][j] {
		return 0
	}
	// test := []int{}
	visited[i][j] = true
	mushrooms := 0
	if forest[i][j] == "C" {
		mushrooms = 1
	}
	fmt.Println("ij ", i, j, "  ", mushrooms, forest[i][j])
	// test = append(test, i, j)
	// fmt.Println("ij ", i, j)
	for nextI := 1; nextI <= 1; nextI++ {
		// fmt.Println("ij3 ", i, j)
		for nextJ := -1; nextJ <= 1; nextJ++ {
			fmt.Println("ij4 ", i, j, "  ", mushrooms)
			fmt.Println("tred ", i+nextI, j+nextJ, "  ")
			mushrooms += pathForest(forest, i+nextI, j+nextJ, visited, test)
			fmt.Println("ij4 ", i, j, "  ", mushrooms)
			*test += fmt.Sprintf("%d %d; ", i, j)
		}
	}
	// 4
	// C.C
	// CWC
	// .WC
	// .WC

	return mushrooms
}

func chackPath(forest [][]string, i, j int) bool {
	if i >= 0 && i < len(forest) && j >= 0 && j < 3 && forest[i][j] != "W" {
		return true
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 5
// ..W
// C.C
// WW.
// CC.
// CWW

func countMushRooms(forest [][]string) (maxMushRooms int) {
	visited := make([][]bool, len(forest))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, 3)
	}
	var test string
	// fmt.Println(visited)
	for j := 0; j < 3; j++ {
		if forest[0][j] != "W" {
			fmt.Println()
			fmt.Println("ppp  ", j, maxMushRooms)
			test += fmt.Sprintf("0 %d; ", j)
			maxMushRooms = max(maxMushRooms, pathForest(forest, 0, j, visited, &test))
			fmt.Println(" Путь: ", test)
			// visited[0][j] = true
			// fmt.Println(maxMushRooms, pathForest(forest, 0, j, visited))
		}
	}

	return
}

// 4
// W.W
// C.C
// WW.
// C..

func main() {
	var x int
	fmt.Scan(&x)

	forest := make([][]string, x)
	str := ""
	forestTmp := make([]string, x)
	for i := 0; i < x; i++ {
		fmt.Scan(&str)
		forestTmp[i] = str
		// fmt.Println(forestTmp[i])
	}
	for i := 0; i < x; i++ {
		forest[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			forest[i][j] = string(forestTmp[i][j])
		}
	}

	fmt.Println(countMushRooms(forest))
}
*/

func main() {
	var n int
	fmt.Scan(&n)

	var in *bufio.Reader
	var out *bufio.Writer
	var a, b int
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	for i := 0; i < n; i++ {

		defer out.Flush()

		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
