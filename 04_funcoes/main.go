package main

import "fmt"

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

func calculosMatematicos(numero1, numero2 int8) (int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("funcao F")
	}

	f()

	resultadoSoma, _ := calculosMatematicos(10, 15)

	fmt.Println(resultadoSoma)
}
