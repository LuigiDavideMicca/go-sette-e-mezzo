package main

import (
	"fmt"
)

func main() {
	fmt.Println("Benvenuto a 7 e Mezzo")
	fmt.Println("---------------------")
	fmt.Println("Puoi chiedere una carta con c oppure fermarti con f")
	fmt.Println("Vuoi giocare? per giocare premi g, per uscire e")
	fmt.Println("---------------------")

	char, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	if char == 'g' {
		playGame()
	} else {
		fmt.Println("Arrivederci")
	}
}
