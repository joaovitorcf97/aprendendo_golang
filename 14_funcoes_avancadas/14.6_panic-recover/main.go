package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar execução")
	}
}

func alunoEstaAprovado(nota1, nota2 float64) bool {
	defer recuperarExecucao()
	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
}
