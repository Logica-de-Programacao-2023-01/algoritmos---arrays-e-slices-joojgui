package main

import "fmt"

//Crie um Slice de inteiros com tamanho 8 e
//solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição.
//Imprima o Slice resultante.

func main() {

	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var indice int
	var indice2 int

	fmt.Println(slice)

	fmt.Println("Informe um indice para trocar o valor de posição :")
	fmt.Scanln(&indice)

	fmt.Println("Informe um indice para trocar o valor de posição :")
	fmt.Scanln(&indice2)

	slice[indice], slice[indice2] = slice[indice2], slice[indice]

	fmt.Println(slice)

}
