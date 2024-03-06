package main

import "fmt"

func funcao01() {
	fmt.Println("Executando função 1")
}

func funcao02() {
	fmt.Println("Executando função 2")
}

func main() {
	// DEFER- ADIAR
	defer funcao01()
	funcao02()
}
