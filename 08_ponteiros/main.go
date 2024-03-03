package main

import "fmt"

func main() {
	var variavel01 int = 10
	var variavel02 int = variavel01

	fmt.Println(variavel01, variavel02)

	variavel01++
	fmt.Println(variavel01, variavel02)

	// PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel03 int
	var ponteiro *int

	variavel03 = 100
	ponteiro = &variavel03
	variavel03 = 200
	fmt.Println(variavel03, *ponteiro)
}
