package main

import (
	"fmt"
	"unicode"
)

// cd .. && mkdir 4_task && cd 4_task && touch main.go && go mod init main.go

func compareString(str1, str2 string) bool {
	size := len(str1)

	for i := 0; i < size; i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true
}

func changeCode(number string) string {
	if len(number) == 7 {
		return "8495" + number
	}
	return number
}

func changednumber(number string) string {
	result := ""
	for i := 0; i < len(number); i++ {
		if number[i] == '+' && number[i+1] == '7' {
			result += "8"
			i = 2
		}
		if unicode.IsDigit(rune(number[i])) {
			result += string(number[i])
		}
	}
	return result
}

func numberFunc(numbers []string) {
	size := len(numbers)
	newNumbers := changednumber(numbers[0])
	// fmt.Println("new ", newNumbers)
	for i := 1; i < size; i++ {
		tmpNumber := changednumber(numbers[i])
		if compareString(changeCode(tmpNumber), changeCode(newNumbers)) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func main() {
	var str string
	var numbers []string
	for i := 0; i < 4; i++ {
		fmt.Scan(&str)
		numbers = append(numbers, str)
	}
	numberFunc(numbers)
}
