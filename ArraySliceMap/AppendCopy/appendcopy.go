package main

import "fmt"

func main() {
	array1 := [3]int{1, 2, 3}
	var slice1 []int

	slice1 = append(slice1, 4, 5, 6)
	fmt.Println(array1, slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	slice1 = append(slice1, 7, 8 , 9)
	slice3 := make([]int, 2, 10)
	copy(slice3, slice1)
	fmt.Println(slice1, slice2, slice3)
	slice3 = append(slice3, 6,7,8,9)
	fmt.Printf("Slice1: %d\nslice3: %d", slice1, slice3)


}