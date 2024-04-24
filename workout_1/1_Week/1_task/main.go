package main

import (
	"os"
	"bufio"
    "fmt"
	"strconv"
)
func condition(tr int, tn int, mode string) int {
	switch mode {
	case "freeze":
		if tr < tn {
			tn = tr
		}
	case "heat":
		if tr > tn {
			tn = tr
		}
	case "auto":

	case "fan":
		tn = tr		
	}

	return tn
}

func main () {
	var tr, tn int
    var mode string
	fmt.Scan(&tr, &tn)
    fmt.Scan(&mode)
    
    result := condition(tr, tn, mode)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	out.WriteString(strconv.Itoa(result))

}