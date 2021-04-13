package main

import (
	"fmt"
	"time"
)

func doisTresQautroX(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base

}

func main() {
	c := make(chan int)

	go doisTresQautroX(8, c)
	s, v := <-c, <-c

	fmt.Println(s, v)
	fmt.Println(<-c)
}
