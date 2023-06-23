package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o tamanho do array: ")
	fmt.Scanln(&n)

	array := make([]int, n)

	fmt.Println("Digite os elementos do array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array[i])
	}

	ordenado := true
	for i := 1; i < n; i++ {
		if array[i] < array[i-1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
