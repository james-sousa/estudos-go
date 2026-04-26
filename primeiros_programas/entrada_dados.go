package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Digite seu nome: ")
	reader.Scan()
	nome := reader.Text()

	fmt.Print("Digite sua idade: ")
	reader.Scan()
	idade := reader.Text()

	fmt.Print("Digite sua altura: ")
	reader.Scan()
	altura := reader.Text()

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
}
