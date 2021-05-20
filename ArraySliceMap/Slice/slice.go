package main

import (
	"fmt"
	"reflect"
)

func main() {

	a1 := [5]int{1, 2, 3, 4, 5} //array
	s1 := []int{1, 2, 3}  //slice
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	s2 := a1[1:4]
	s3 := a1[3:4]
	fmt.Println(s2, s3)

	s4 := s2[:1] //slice de um slice
	s5 := a1[:1] //slice da primeira posição de um array 
	s6 := a1[1:] //slice do indice 1 até o último indice de array de 5 posições


	fmt.Println(s4,s5,s6)


}
