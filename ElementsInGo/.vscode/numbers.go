package main

import "fmt"

func main() {
	sum := 0
	n := 10
	//fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Print("Sum: ", sum)
}
