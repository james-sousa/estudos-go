package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Println("Erro: Divisão por zero")
		return 0, fmt.Errorf("divisão por zero")
	}
	return a / b, nil
}

func somaVariadica(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	resultadoSoma := soma(5, 3)
	fmt.Println("Resultado da soma:", resultadoSoma)

	resultadoDivisao, err := divisao(10, 2)
	if err != nil {
		fmt.Println("Erro na divisão:", err)
	} else {
		fmt.Println("Resultado da divisão:", resultadoDivisao)
	}

	resultadoSomaVariadica := somaVariadica(1, 2, 3, 4, 5)
	fmt.Println("Resultado da soma variádica:", resultadoSomaVariadica)
}