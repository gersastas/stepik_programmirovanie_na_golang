package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a > 0:
		fmt.Println("Число положительное")
	case a < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}
