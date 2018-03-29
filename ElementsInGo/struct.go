package main
import "fmt"
import "math"

type Circle struct {
	x,y,r float64
}
//Function
func circleArea(c *Circle) float64 {
	return math.Pi*c.r*c.r
}
//Method
func (c *Circle) area() float64 {
	return math.Pi*c.r*c.r
}
func main(){
	//c := new (Circle)
	c := Circle{x:0,y:0,r:4}
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())
}