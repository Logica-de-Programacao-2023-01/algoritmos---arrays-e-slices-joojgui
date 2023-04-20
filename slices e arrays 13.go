package main

import "fmt"

func main() {
	//Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um
	//número que será adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.

	Array := [7]int{1, 2, 3, 4, 5, 6, 7}

	var numero int

	fmt.Println(Array)

	fmt.Println("Insira um número para somar a o primeiro e o ultimo número do array :")
	fmt.Scan(&numero)

	Array[0] = 1 + numero

	Array[6] = 7 + numero

	fmt.Println(Array)

}
