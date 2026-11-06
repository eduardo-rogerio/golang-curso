package main

import "fmt"

func main() {
	result := sum(23, 45)
	fmt.Println("o resultado da soma é: ", result)

  q, r := divide(10, 2)
  fmt.Printf("o resultado da divisão é %d e o resto da divisão é %d\n", q, r)

  multiplicationTable()
}

func sum(a, b int8) int8 {
	return a + b
}

// função com multiplos retornos
func divide(a, b int8) (int8, int8) {
  quotient := a / b
  remainder := a % b
  return quotient, remainder
}

// função sem retorno
func multiplicationTable() {
  for i := 0; i < 11; i++ {
    for j := 0; j < 11; j++ {
      r := i * j
    fmt.Printf("%d x %d = %d\n", i, j, r)
    }
    fmt.Println("------------------------")
  }
}
