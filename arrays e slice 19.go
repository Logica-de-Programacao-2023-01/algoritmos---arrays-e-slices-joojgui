package main

import "fmt"

//Faça um programa que leia dois arrays de inteiros de tamanho ne gere um terceiro array que seja a soma dos dois primeiro

func main() {

	array1 := []int{1, 2, 3, 4, 5}
	array2 := []int{6, 7, 8, 9, 10}

	result := make([]int, len(array1))

	for i := 0; i < len(array1); i++ {
		result[i] = array1[i] + array2[i]
	}

	fmt.Println(result)
}
