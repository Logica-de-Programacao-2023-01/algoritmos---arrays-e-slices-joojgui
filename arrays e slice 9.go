package main

import "fmt"

//Crie um Array de floats com 6 elementos.
//Solicite ao usuário que informe um número que será adicionado em todas as posições do Array.
//]Imprima o Array resultante

func main() {

	var Array = [6]float64{4.9, 6.5, 2.3, 1.9, 8.9, 7.5}
	var valor float64

	fmt.Print("Informe um número: ")
	fmt.Scan(&valor)

	for i := 0; i < len(Array); i++ {
		Array[i] += valor
	}

	fmt.Println("A array agora fica : ", Array)
}
