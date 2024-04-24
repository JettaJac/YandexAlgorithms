package main2

import (
	"fmt"
)

// cd .. && mkdir 6_task && cd 6_task && touch main.go && go mod init main.go
// тест 89 20 45 1 11
// 1 1 2 1 1		1 1
// 2 1 1 1 1		0 1
// 89 20 44 1 11 		2 3
// 89 20 40 1 11	-1 -1
// 1 1 10 2 1		1 1
// 7 1 10 2 1		0 1
// 8 4 10 2 3		-1 -1
// 12 6 14 2 3		-1 -1
// 7 4 11 1 2		1 2
// 15 5 20 3 4		-1   -1
// 12 3 18 2 3		2 1
// 9 3 18 2 3		1 3
// 10 3 18 2 3		2 1
// 100 1 2 1 1		0 1
// 100 2 2 1 1
// 100 2 2 1 1      0 0
// 2 1 2 1 1        1 1
// 45 2 45 2 3      -1 -1
// 43 20 43 1 11
// 18 3 18 2 3

func checkInfo(k2, p2, n2, nAF, m int) bool {
	// fmt.Println("checkInfo ", k2)
	// fmt.Println(p2*m*nAF, "  ", m*(p2-1), "  kv", k2, "  ", n2*nAF+(p2-1)*m*nAF, "  ", (n2-1)*nAF+(p2-1)*m*nAF)
	if /*!(k2 < p2*m*nAF && k2 > m*(p2-1)*nAF) || !*/ k2 <= n2*nAF+(p2-1)*m*nAF && k2 > (n2-1)*nAF+(p2-1)*m*nAF {
		// fmt.Println("True")
		return true
	}

	return false
}

func numberApartFloor(m, k2, n2, p2 int) int {

	tmpInt := k2 / ((p2-1)*m + n2)
	// fmt.Println(tmpInt)
	if k2 > tmpInt*n2+(p2-1)*m*tmpInt {
		return tmpInt + 1
	}
	return tmpInt
}

func nuberAppartEntrance(nAF, m int) int {
	return nAF * m
}

func BaseFunc(k1, m, k2, p2, n2 int) string {
	var p1, n1 int
	nAF := numberApartFloor(m, k2, n2, p2)
	// fmt.Println("Количество квартир на этаже_ ", nAF)

	if checkInfo(k2, p2, n2, nAF, m) && k1 != 0 && n2 <= m {
		// Количество квартир в подъезде
		nAE := nuberAppartEntrance(nAF, m)

		// Подъезд
		p1 = (k1 / nAE)
		if k1 > p1*nAE {
			p1 += 1
		}

		// Этаж
		n1 = (k1 % nAE) / nAF
		// fmt.Println(k1 > nAF*n1+(p1-1)*nAE, " | ", k1, " ", nAF*n1+(p1-1)*nAE)
		if k1%nAE == 0 {
			n1 = m
		} else if k1 > nAF*n1+(p1-1)*nAE {
			// fmt.Println("ggg ", k1, n1, p1, " ", nAF*n1+(p1-1)*nAE)
			n1 += 1
		}
		// fmt.Println("Количество квартир в подъезде_ ", nAE)
	} else {
		p1 = -1
		n1 = -1
	}
	// }

	if n2 == 1 && (p2 == 1 || m == 1) && k1 != 0 {
		// fmt.Println("ggg")
		n1 = 0
		p1 = 0
		if k1 <= k2 && p2 == 1 || k1 == 1 {
			p1 = 1
		}
		if m == 1 {
			n1 = 1
		}
	}
	// fmt.Println(p1, " ggg", n1)
	return (fmt.Sprintf("%d %d", p1, n1))
}

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)

	// if n2 != 1 && m != 1 {
	// Количество квартир на этаже

	fmt.Println(BaseFunc(k1, m, k2, p2, n2))
}
