package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Jose Claudio":      22390.99,
		"Vitor Grabriel":    15000.99,
		"Yanzao do paredão": 15000,
	}

	funcsESalarios["Jack Terror"] = 1999.99

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
