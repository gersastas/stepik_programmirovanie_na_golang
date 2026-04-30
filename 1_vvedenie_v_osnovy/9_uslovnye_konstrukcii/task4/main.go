package main

import "fmt"

const (
	ten             = 10
	hundred         = 100
	thousand        = 1000
	tenThousand     = 10000
	hundredThousand = 100000
)

func main() {
	var num int
	fmt.Scan(&num)

	if num < hundredThousand || num > 999999 {
		fmt.Println("NO")
		return
	}

	numOne := num % ten
	numTwo := (num / ten) % ten
	numTree := (num / hundred) % ten
	numFour := (num / thousand) % ten
	numFive := (num / tenThousand) % ten
	numSix := num / hundredThousand

	if numOne+numTwo+numTree == numFour+numFive+numSix {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
