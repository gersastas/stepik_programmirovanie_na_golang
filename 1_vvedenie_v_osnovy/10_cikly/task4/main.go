package main

import "fmt"

func main() {
	max := 0
	count := 0

	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		if x > max {
			max = x
			count = 1
		} else if x == max {
			count++
		}
	}

	fmt.Println(count)
}
