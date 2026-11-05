package main

import "fmt"

func main() {
	// loops em golang
	// for comum
	for i := 0; i < 10; i++ {
		fmt.Println("valor de i: ", i)
	}

	// for como while
	i := 10
	for i < 20 {
		fmt.Println("valor de i:", i)
		i++
	}

  // for como loop infinito
  for {
    fmt.Println("loop infinito...")
    break // para o qualquer loop
  }

  // range de slice
  number := []int{10, 20, 30}
  for _,value := range number {
    fmt.Printf("valor: %d\n", value)
  }

  // range de map
  capitals := map[string]string {
    "Brasil": "Brasília",
    "França": "Paris",
    "Japão": "Tokio",
  }

  for country, capital := range capitals {
    fmt.Printf("A capital do %s é %s\n", country, capital)
  }

}
