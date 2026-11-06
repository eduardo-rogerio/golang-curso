package main

import "fmt"

func main() {
  // clousure simples
  closure := func(a, b int) int {
    return a + b
  }

  fmt.Println(closure(45, 33))
  odd := oddGenerator()
  fmt.Println(odd()) 
  fmt.Println(odd()) 

  fmt.Println(max(34,230,33,2))
}

// função que retorna closure
func oddGenerator() func() int {
  i := -1
  return func() int {
    i += 2
    return i
  }
}

// função variádica 
func max(numbers ...int) int {
  if len(numbers) == 0 {
    return 0
  }

  bigger := numbers[0]
  for _, number := range numbers {
    if number > bigger {
      bigger = number
    }
  }

  return bigger
}
