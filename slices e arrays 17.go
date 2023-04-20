package main

//Crie um Array de inteiros com 10 elementos.
//Calcule e imprima a soma dos elementos nas posições pares do Array.

import "fmt"

func main() {

	var array = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := 0

	for i := 0; i < len(array); i += 2 {
		sum += array[i]
	}

	fmt.Println(array)

	fmt.Println(sum)

}




