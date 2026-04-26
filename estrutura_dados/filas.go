package main

import "fmt"

type Fila []int

func (f *Fila) Enqueue(valor int) {
	*f = append(*f, valor)
}

func (f *Fila) SairFila() int {
	valor := (*f)[0]
	*f = (*f)[1:]
	return valor
}

func main() {
	fila := Fila{}

	fila.Enqueue(1)
	fila.Enqueue(2)
	fila.Enqueue(3)

	fmt.Println("Fila:", fila)

	valorRemovido := fila.SairFila()
	fmt.Println("Valor removido:", valorRemovido)
	fmt.Println("Fila após remoção:", fila)
}
