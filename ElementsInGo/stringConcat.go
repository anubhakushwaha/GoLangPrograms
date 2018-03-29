package main

import "fmt"

func main(){
	fmt.Print("Enter the first string: ")
	var first string
	fmt.Scanln(&first)
	fmt.Print("Enter the second string: ")
	var second string
	fmt.Scanln(&second)
	fmt.Print(first+second)
}