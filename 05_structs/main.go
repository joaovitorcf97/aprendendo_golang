package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "João"
	u.idade = 27
	fmt.Println(u)

	endereco1 := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Ana", 22, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 21}
	fmt.Println(usuario3)
}
