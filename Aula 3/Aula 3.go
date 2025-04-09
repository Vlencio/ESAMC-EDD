package main

import "fmt"

func exemploAula3 () {
  var a int = 42
  var p *int = &a
  fmt.Println("Valor de a:", a)
  fmt.Println("Endereço de a:", &a)
  fmt.Println("Valor de p (endereço de a):", p)
  fmt.Println("Valor apontado por p:", *p)
  fmt.Println("Valor de a antes da modificação:", a)
  *p = 100
  fmt.Println("Valor de a após a modificação:", a)
}

func exec1_ponteiro() {
  var b float64 = 0
  var p2 *float64 = &b

  fmt.Println("Valor antes de mudar: ", b)
  *p2 = 69
  fmt.Println("Valor depois de mudar: ", b)
}

func ponteiro_swap(a, b *int){
  fmt.Println(*a, *b)
  temp := *a
  *a = *b
  *b = temp
  fmt.Println(*a, *b)
}

