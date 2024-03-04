package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		time.Sleep(time.Second)
		fmt.Println(i)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Lucas", "Pedro"}

	for _, nome := range nomes {
		fmt.Println(nome)
		time.Sleep(time.Second)
	}

}
