package main

import "fmt"

func operacao(x int, y int) {
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)

	if y != 0 {
		fmt.Println(x / y)
	} else {
		fmt.Println("erro")
	}
}

func main() {
	// Problemas:
	// 1. Tudo na função main
	// 2. Nomes não descritivos
	// 3. Sem tratamento de erros
	// 4. Lógica repetida
	// 5. Formatação inconsistente
	x := 10
	y := 5

	operacao(x, y)

}
