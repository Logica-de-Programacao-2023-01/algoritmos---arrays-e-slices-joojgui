package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	var w float64
	var produto float64

	fmt.Print("insira 4 números decimais: ")

	fmt.Scan(&x, &y, &z, &w)

	numeros := [4]float64{x, y, z, w}
	produto = 1

	for _, numeros := range numeros {

		produto *= numeros

	}

	fmt.Println(produto)

}

