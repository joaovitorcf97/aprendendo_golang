package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	pessoa1 := pessoa{"João", 27, 170}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Cienda da Computação", "USP"}
	fmt.Println(estudante1)
}
