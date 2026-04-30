package main

import "fmt"

const (
	minNumber = 10
	maxNumber = 100
)

func main() {
	for {
		var x int
		fmt.Scan(&x)

		if x < minNumber {
			continue
		}

		if x > maxNumber {
			break
		}

		fmt.Println(x)
	}
}
