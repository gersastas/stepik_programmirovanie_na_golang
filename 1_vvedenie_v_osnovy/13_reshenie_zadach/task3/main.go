package main

import "fmt"

const maxSize = 86399

func main() {
	var k int
	for {
		fmt.Scan(&k)
		if k > 0 && k < maxSize {
			break
		}
	}
	hour := k / 3600
	rest := k - hour*3600
	minute := rest / 60

	fmt.Printf("It is %d hours %d minutes.", hour, minute)
}
