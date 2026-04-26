package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
	Email string
}

func (p Pessoa) Apresentar() string {
	return fmt.Sprintf("Olá, meu nome é %s, tenho %d anos e meu email é %s.", p.Nome, p.Idade, p.Email)
}

func (p *Pessoa) Aniversario() {
	p.Idade++
}

type Funcionario struct {
	Pessoa
	Cargo   string
	Salario float64
}

func main() {
	p1 := Pessoa{Nome: "Gustavo", Idade: 10, Email: "gustavo@example.com"}
	fmt.Println(p1.Apresentar())
	p1.Aniversario()
	fmt.Println("Após aniversário:")
	fmt.Println(p1.Apresentar())

	f1 := Funcionario{Pessoa: Pessoa{Nome: "Maria", Idade: 30, Email: "maria@example.com"}, Cargo: "Engenheira", Salario: 5000.0}
	fmt.Println(f1.Apresentar())
	fmt.Printf("Cargo: %s, Salário: %.2f\n", f1.Cargo, f1.Salario)
}
