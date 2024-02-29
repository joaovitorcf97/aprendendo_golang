package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	var numero02 uint32 = 1000
	fmt.Println(numero02)

	var numero03 rune = 12345
	fmt.Println(numero03)

	var numero04 byte = 123
	fmt.Println(numero04)

	var numero05 float32 = 1234.56
	fmt.Println(numero05)

	var str string = "Texto"
	fmt.Println(str)

	var booleano01 bool = true
	fmt.Println(booleano01)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
