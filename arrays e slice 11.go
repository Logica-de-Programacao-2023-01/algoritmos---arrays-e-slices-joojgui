package main

import (
	"fmt"
)

//Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida,
//solicite ao usuário que informe um índice de linha e outro de coluna e
//imprima o valor armazenado nessa posição da matriz.

func main() {
	var coluna int
	array := [2][3]int{{1, 2, 3}, {6, 7, 8}}
	var linha int

	fmt.Println("Informe um índice de linha (0 e 1 ), e um de coluna ( 0 a 2 ):")
	fmt.Scan(&linha, &coluna)

	fmt.Println("A posição do array é :", array[linha][coluna])

}
