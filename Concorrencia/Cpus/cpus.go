package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Quantidade Nucleos da sua CPU:", runtime.NumCPU())
}
