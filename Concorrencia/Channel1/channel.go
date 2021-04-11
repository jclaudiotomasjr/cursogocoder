package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 8 // enviando dados para o canal escrita
	<-ch    // recebendo dados do canal leitura

	ch <- 9

	fmt.Println(<-ch)

}
