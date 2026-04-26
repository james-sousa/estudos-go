package main

import "fmt"

func main() {
	var arrays [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arrays)

	frutas := []string{"maçã", "banana", "laranja"}
	frutas = append(frutas, "uva")

	for i, v := range frutas {
		fmt.Printf("Índice: %d, Valor: %s\n", i, v)
	}

	idades := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 28,
	}

	idades["Dave"] = 35

	if idade, ok := idades["Alice"]; ok {
		fmt.Printf("Idade de Alice: %d\n", idade)
	}

	for nome, idade := range idades {
		fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
	}

	delete(idades, "Bob")

	fmt.Println("Idades após remoção:", idades)
}
