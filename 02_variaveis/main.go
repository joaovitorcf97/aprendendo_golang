package main

import "fmt"

func main() {
	var variavel01 string = "Variavel 01"
	variavel02 := "Variavel 02"

	fmt.Println(variavel01)
	fmt.Println(variavel02)

	var (
		variavel03 string = "Variavel 03"
		variavel04 string = "Variavel 04"
	)

	fmt.Println(variavel03, variavel04)

	const constante01 string = "Constante 01"

	fmt.Println(constante01)
}
