package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobreNome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobreNome
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	var coisa imprimivel = pessoa{"Jose Claudio", "Tomás"}
	fmt.Println(coisa.toString())

	imprimir(coisa)

	coisa = produto{"Camisa do Botafogo", 100.00}
	fmt.Println(coisa.toString())

	p2 := produto{"Tenis", 300.00}
	imprimir(p2)
}
