package main

import "fmt"

type Pilha []int

func (p *Pilha) Push(valor int) {
	*p = append(*p, valor)
}

func (p *Pilha) Pop() int {
	indice := len(*p) - 1
	valor := (*p)[indice]
	*p = (*p)[:indice]
	return valor
}

func main() {
	pilha := Pilha{}

	pilha.Push(1)
	pilha.Push(2)
	pilha.Push(3)

	fmt.Println("Pilha:", pilha)

	valorRemovido := pilha.Pop()
	fmt.Println("Valor removido:", valorRemovido)
	fmt.Println("Pilha após remoção:", pilha)
}
