package main

import "fmt"

type curso struct {
	nome string
}

func main() {

	var coisa interface{}
	fmt.Println(coisa)

	coisa = "Jujuba"
	fmt.Println(coisa)

	coisa = 500.99
	fmt.Println(coisa)

	type dinamico interface{}

	var coisa2 dinamico = "Vitor"
	fmt.Println(coisa2)

	coisa2 = "Stefania"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	type dinamico2 interface{}
	var inteiro dinamico2 = 8899
	fmt.Println(inteiro)

	inteiro = curso{"Golang explorando a linguagem do Google"}
	fmt.Println(inteiro)
}
