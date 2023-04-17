package main

import "fmt"

//Crie um Array de inteiros com 10 elementos. Em seguida,
//solicite ao usuário que informe um valor e verifique se esse valor está presente no Array.
//Imprima o resultado da busca

func main() {
	array := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	var valor int
	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)

	encontrado := false
	for _, elemento := range array {
		if elemento == valor {
			encontrado = true
			break
		}
	}
	if encontrado {
		fmt.Printf("O valor %d está presente no array.", valor)
	} else {
		fmt.Printf("O valor %d não está presente no array.", valor)
	}

}
