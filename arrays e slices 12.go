package main

import "fmt"

// rie um Array de inteiros com 5 elementos. Em
// seguida, crie um novo Slice que atendeu apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
func main() {

	Array := [5]int{6, 2, 15, 21, 7}

	var NovoArray []int

	for _, numero := range Array {

		if numero%3 == 0 {

			NovoArray = append(NovoArray, numero)

		}

		fmt.Println(NovoArray)

	}

}
