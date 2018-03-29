package main
import "fmt"

var x string = "Hello World"
func main() {
	// Go is lexically scoped using blocks
	// this will work but if x were local it would compile error
	x := 5
	fmt.Println(x)
}