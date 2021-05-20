package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 14
	fmt.Println(s)

	s = make([]int, 10, 20)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 2, 3, 4, 8, 9, 11, 14, 16, 20, 22)
	fmt.Println(s, len(s), cap(s))
}
