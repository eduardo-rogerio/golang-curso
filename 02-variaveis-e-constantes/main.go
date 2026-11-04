package main

import "fmt"

const Teste1 = "constant teste"

func main(){
  // usando o var
  var name1 string
  name1 = "usando var: teste"

  // fazendo atribuição direta
  var name2 string = "iniciando valor com var: Pedro"

  // inferencia de tipo
  var name3 = "inferencia de tipo: Pedro"

  // declaração curta
  name4 := "declaração curta: Pedro"

  fmt.Println(name1)
  fmt.Println(name2)
  fmt.Println(name3)
  fmt.Println(name4)

  // tipos booleanos só aceitam os tipos true ou false. Não aceitam 1 e 0
  var isEnable bool = false
  fmt.Println(isEnable)

  // declarando varias variáveis em uma única linha

  var a, b, c = "a", 50, true

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}
