package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	slice := make([]int, n)

	for i := 0; i < len(slice); i++ {
		fmt.Scan(&slice[i])
	}
	fmt.Println(slice[3])

}
