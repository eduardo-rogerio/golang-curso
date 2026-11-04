package main

import "fmt"

func main() {
	// switch case em golang
	day := "Quintta-feira"

	switch day {
	case "Segunda-feira", "Terça-feira", "Quata-feira", "Quinta-feira", "Sexta-feira", "Sábado":
		fmt.Println("dia útil")
	case "Domingo":
		fmt.Println("não é dia últil")
	default:
		fmt.Println("dia não reconhecido")
	}
}
