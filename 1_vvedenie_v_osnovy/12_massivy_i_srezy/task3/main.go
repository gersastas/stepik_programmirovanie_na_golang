package main

import "fmt"

func main() {
	arr := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		arr[i] = a
	}

	max := arr[0]
	for i := 1; i < 5; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println(max)
}
