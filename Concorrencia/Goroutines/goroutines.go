package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//go fale("Maria", "Ei...", 500)
	//go fale("João", "Opa...", 500)
		
	go fale("Stefania", "Entendi!!", 14)
	fale("Vitor", "Parabéns", 8)

}
