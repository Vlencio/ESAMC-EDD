package main

import (
	"fmt"
)

func vetor() {
	vetor_RA := [6]int{1, 2, 4, 1, 3, 5}
	for _, valor := range vetor_RA {
		fmt.Print(valor)
	}

	fmt.Print(vetor_RA[0])
	fmt.Print(vetor_RA[1])
	fmt.Print(vetor_RA[2])
	fmt.Print(vetor_RA[3])
	fmt.Print(vetor_RA[4])
	fmt.Print(vetor_RA[5])
	for i := 0; i < 6; i++ {
		fmt.Print(vetor_RA[i])
	}

}

func matriz() {
	var matriz [3][3]string

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				fmt.Println(i, j)
				fmt.Print("Digite o indice do aluno ", i+1, ": ")
				fmt.Scan(&matriz[i][j])
			}

			if j == 1 {
				fmt.Println(i, j)
				fmt.Print("Digite o RA do aluno ", i+1, ": ")
				fmt.Scan(&matriz[i][j])
			}
			if j == 2 {
				fmt.Println(i, j)
				fmt.Print("Digite o nome do aluno ", i+1, ": ")
				fmt.Scan(&matriz[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("Dados dos Alunos:")
	fmt.Println("-----------------")
	fmt.Println("Indice  RA  Nome")
	for i, _ := range matriz {
		for j, _ := range matriz {
			fmt.Print(matriz[i][j])
			fmt.Print("    ")
		}
		fmt.Println()
	}
}

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func execAula2() {
	var num int

	fmt.Println("Digite quantos produtos deseja adicionar: ")
  fmt.Scan(&num)

  
  
  for i := 0; i < num; i++ {
    
  }

}
