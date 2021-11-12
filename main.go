package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	escolha := menu()

	var (
		min, max, quant int
	)

	switch escolha {

	case 1:
		min = minimo()
		max = maximo()
		quant = quantidade()

	case 2:
		min = 1
		max = 100
		quant = quantidade()

	}

	fmt.Println(numeroAleatorio(min, max, quant))

}

func menu() int {

	var escolha int

	fmt.Println("Escolha o modo de Geração")
	fmt.Println("1 - Manual(escolha o minimo e o maximo)")
	fmt.Println("2 - Automático(gera de 1 a 100)")
	fmt.Print("Opção: ")
	fmt.Scan(&escolha)

	return escolha
}

func quantidade() int {
	fmt.Print("Quantos números aleatórios você quer gerar? ")
	var quantidade int
	fmt.Scan(&quantidade)
	return quantidade
}

func numeroAleatorio(min, max, quant int) []int {

	var numsAl []int
	rand.Seed(time.Now().UnixNano())

	for {

		if quant != 0 {

			numAl := rand.Intn(max-min) + min
			numsAl = append(numsAl, numAl)
			quant--
		} else {
			break
		}
	}

	return numsAl

}

func minimo() int {
	fmt.Print("Qual o número mínimo? ")
	var min int
	fmt.Scan(&min)
	return min
}

func maximo() int {
	fmt.Print("Qual o número maximo? ")
	var max int
	fmt.Scan(&max)
	return max
}
