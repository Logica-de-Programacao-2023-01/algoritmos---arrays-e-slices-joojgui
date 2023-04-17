package main

import "fmt"

//Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor.
//Remova todas as ocorrências desse valor no Slice e imprima o resultado.

func main() {

	var Strings = []string{"João", "Maria", "jose", "Matheus", "eduardo", "pedro", "camila", "carlos"}
	var Valor string

	fmt.Print("Insira um nome :")
	fmt.Scanln(&Valor)

	Slice := []string{}

	for _, Valor2 := range Strings {

		if Valor2 != Valor {

			Slice = append(Slice, Valor2)
		}
	}

	fmt.Println("A nova slice fica: ", Slice)

}
