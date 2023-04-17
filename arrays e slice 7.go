package main

import "fmt"

//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

func main() {

	var Slice = []int{5, 8, 3, 2, 4}
	var valor int

	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)

	encontrado := false
	for _, elemento := range Slice {
		if elemento == valor {
			encontrado = true
			break
		}
	}

	if !encontrado {
		Slice = append(Slice, valor)

		fmt.Println("O valor inserido na slice fica :", Slice)
	} else {

		fmt.Printf("O valor %d ja esta presenta na slice ", valor)

	}

}
