package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	hundred := (num / 100) % 10
	ten := (num / 10) % 10
	unit := num % 10

	if hundred != ten && ten != unit && unit != hundred {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
