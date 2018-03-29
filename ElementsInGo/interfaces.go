package main
import "fmt"
import "math"

type Shape interface {
	area() float64
}

type Circle struct {
	x,y,r float64
}
func (c *Circle) area() float64 {
	return math.Pi*c.r*c.r
}

type Rectangle struct {
	x,y float64
}
func (r *Rectangle) area() float64 {
	return r.x*r.y
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for i:=0; i<len(shapes);i++ {
		area += shapes[i].area()
	}
	return area
}
func main(){
	c := Circle{x:0, y:0, r:5}
	r := Rectangle{x:2, y:3}

	fmt.Println(totalArea(&c,&r))
}