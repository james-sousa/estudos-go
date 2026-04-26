package main

import "fmt"

type No struct {
	Valor    int
	Esquerda *No
	Direita  *No
}

func (n *No) Inseir(valor int) {
	if valor < n.Valor {
		if n.Esquerda == nil {
			n.Esquerda = &No{Valor: valor}
		} else {
			n.Esquerda.Inseir(valor)
		}
	} else {
		if n.Direita == nil {
			n.Direita = &No{Valor: valor}
		} else {
			n.Direita.Inseir(valor)
		}
	}
}

func (n *No) Buscar(valor int) bool {
	if n == nil {
		return false
	}
	if valor == n.Valor {
		return true
	} else if valor < n.Valor {
		return n.Esquerda.Buscar(valor)
	} else {
		return n.Direita.Buscar(valor)
	}
}

func main() {
	raiz := &No{Valor: 10}
	raiz.Inseir(5)
	raiz.Inseir(15)
	raiz.Inseir(3)
	raiz.Inseir(7)
	fmt.Printf("imprimir arvore")
	fmt.Printf("Raiz: %d\n", raiz.Valor)
	fmt.Printf("Esquerda: %d\n", raiz.Esquerda.Valor)
	fmt.Printf("Direita: %d\n", raiz.Direita.Valor)
	fmt.Println("Buscar 7:", raiz.Buscar(7)) // true
	fmt.Println("Buscar 4:", raiz.Buscar(4)) // false
}
