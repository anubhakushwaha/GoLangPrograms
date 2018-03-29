package main
import "fmt"
func main(){
	fmt.Println(len("Hello World"))
	//Prints the character at first index of the word
	fmt.Println("Hello World"[1])
	fmt.Println("Hello"+"World")

	var x string = "hello"
	var y string = "hello"
	// will show true
	fmt.Println(x==y)

	z := "Hello World"

	// Will show error as z has been linked as a string
	z := 5

	fmt.Println(z)
}