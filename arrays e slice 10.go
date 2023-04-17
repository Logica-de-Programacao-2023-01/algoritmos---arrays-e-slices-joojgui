package main

import (
	"fmt"
)

//Crie um Slice de inteiros com tamanho 10
//e imprima o valor mínimo e o valor máximo armazenados no Slice.

func main() {

	var Slice = []int{4, 7, 3, 10, 3, 6, 8, 9, 13}

	min := 999999

	max := -99999

	for _, x := range Slice {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}

	fmt.Println(Slice)

	fmt.Println("O maior valor é", max)

	fmt.Println("O menor valor é", min)

}
