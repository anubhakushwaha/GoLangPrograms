package main
import "fmt"

func funAll(){
	friends := map[string]map[string]string{
		"anubha": map[string]string{
			"class":"XII",
			"school":"RLB",
		},
		"archit": map[string]string{
			"class":"XII",
			"college":"RLB",
		},
	}

	if name, ok := friends["anubha"]; ok{
		fmt.Println(name["school"])
	}
}
func main(){

	//<------ arrays------>
	var x[5]float64
	// initialized whole array with 0
	// [0 0 0 0 0]

	x = [5]float64{33,44,35,11,2}
	//x[4]=100

	//Error as mismatched datatypes float and int
	//fmt.Print(x[4]/len(x))

	fmt.Print(x[4]/float64(len(x)))

	//<------ slices------>
	//unlike arrays their length can change

	y :=make([]float64, 5, 10)

	arr := [5]float64{1,2,3,4,5}
	z := arr[0:5]

	fmt.Println(y,z)

	//two built-in functions
	//append--copy
	slice1 := []int{1,2,3,4,5}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	//[1 2 3 4 5] [1 2 3 4 5 4 5]

	slice3 := make([]int, 3)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)
	//[1 2 3 4 5] [1 2 3]

	//<------ maps -------->
	// var mp map[string]int
	// //map of strings to int's
	// mp["key"]=10
	// fmt.Println(mp["key"])
	// //gives runtime error

	sample := make(map[string]int)
	sample["key"] = 10
	fmt.Println(sample["key"])

	// delete(sample, "key")
	// fmt.Println(sample["key"])

	/* When key not present returns zero value for the value type*/
	if name, ok := sample["key"]; ok{
		fmt.Println(name, ok)
	}

	funAll();
}