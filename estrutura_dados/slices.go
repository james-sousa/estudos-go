package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	// Adicionando um elemento ao slice
	numeros = append(numeros, 6)
	fmt.Println(numeros)
	var novoNumero int
	fmt.Printf("Adicionando: ")
	fmt.Scanln(&novoNumero)
	numeros = append(numeros, novoNumero)
	fmt.Println(numeros)
}
