package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkFuncS(a []int) bool {
	tmp := a[0]
	for i := 1; i < len(a); i++ {
		fmt.Println(i, a[i], len(a), tmp)
		if tmp >= a[i] {
			return false
		}
		tmp = a[i]
	}
	return true
}

func checkFunc(a, b, c int) bool {

	if a < b && b < c {
		return true
	}
	return false

}

func readStdinInt(input *os.File) []int {
	res := []int{}

	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		res := append(res, 2, 1)
		return res
	}
	data = strings.TrimSpace(data)
	resStr := strings.Split(data, " ")
	// fmt.Println(resStr)

	for _, i := range resStr {
		n, err := strconv.Atoi(i)
		if err != nil {
			res = append(res, 2, 1)
			return res
		}
		res = append(res, n)
	}

	return res
}

func main() {
	// var a, b, c int
	// fmt.Scan(&a, &b, &c)

	// readStdinInt(os.Stdin)
	// if checkFunc(a, b, c) {
	if checkFuncS(readStdinInt(os.Stdin)) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")

}
