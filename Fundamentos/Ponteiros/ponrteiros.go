package main

import "fmt"

func main() {
	j := 8

	var s *int = nil

	fmt.Println(j, s)

	s = &j
	*s++

	fmt.Println(s, &j, j, *s)

}
