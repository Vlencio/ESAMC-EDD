package main

import (
	"fmt"
)

func ParOuImpar(x int) string {
	if x%2 == 0 {
		return "Par"
	}
	return "Impar"
}

func SomaAte(n int) int {
	var soma int = 0
	for i := 1; i <= n; i++ {
		fmt.Println(soma)
		soma += i
		fmt.Println(soma)
	}
	return soma
}

func ClassificarNota(nota int) string {
	switch {
	case nota >= 9:
		return "Excelente"
	case nota >= 7:
		return "Bom"
	case nota >= 5:
		return "Regular"
	case nota >= 3:
		return "Ruim"
	}
	return "Muito Ruim"
}
