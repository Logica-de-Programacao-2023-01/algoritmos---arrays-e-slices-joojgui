package main

import "fmt"

//Crie um Array de floats com 10 elementos.
//Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5.
//Imprima o novo Slice.

func main() {

	var array = [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var slice = []float64{}

	for _, numero := range array {

		if numero > 5 {

			slice = append(slice, numero)

		}

	}

	fmt.Println(slice)

}
