package main

import "fmt"

//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida,
//solicite ao usuário que informe um índice de linha e outro de coluna e
//imprima o valor armazenado nessa posição da matriz.

func main() {
	var linha int
	var coluna int
	array := [2][3]int{{1, 2, 6}, {2, 5, 10}}

	