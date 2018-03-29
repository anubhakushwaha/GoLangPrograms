package main
import "fmt"

func main(){
	x := []int{34,12,45,44,13,55,75,}
	ans :=100
	for i:=0;i<len(x);i++ {
		if ans>x[i] {
			ans=x[i]
		}
	}
	fmt.Println(ans)
}