package main

import (
	"fmt"
	"modulo/gocoder/html"
	"time"
)

func flash(url1, url2, url3 string) string {
	c1 := html.Titulo(url1)
	c2 := html.Titulo(url2)
	c3 := html.Titulo(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1500 * time.Millisecond):
		return "Ninguém venceu!"
		//default:
		//	return "Sem resposta!"
	}
}
func main() {
	champ := flash(
		"https://www.youtube.com/",
		"https://www.google.com.br/",
		"https://golang.org/",
	)
	fmt.Println(champ)
}
