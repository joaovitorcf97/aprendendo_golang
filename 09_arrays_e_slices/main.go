package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]string

	array1[0] = "Posição 01"
	fmt.Println(array1)

	array2 := [4]string{"Posição 01", "Posição 02", "Posição 03", "Posição 04"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	slice := []int{10, 20, 30}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 40)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[2] = "Posiçãp alterada"
	fmt.Println(array2)

	// Arrays internos
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacidade
}
