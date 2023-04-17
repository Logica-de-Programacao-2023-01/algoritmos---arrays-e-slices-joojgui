package main

import "fmt"

//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os
//valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados

func main() {

	var tamanho int
	fmt.Print("Informe o tamanho do slice: ")
	fmt.Scanln(&tamanho)

	slice := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("Informe o valor do elemento %d: ", i)
		fmt.Scanln(&slice[i])
	}

	soma := 0
	for _, valor := range slice {
		soma += valor
	}

	fmt.Println("Soma dos valores: ", soma)
}
