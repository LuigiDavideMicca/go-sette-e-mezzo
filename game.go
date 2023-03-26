package main

import (
	"fmt"
	"os"
)

func playGame() {
	fmt.Println("Il gioco ha inizio")
	fmt.Println("---------------------")
	fmt.Println("La Matta è il Re di Cuori")
	var score float32
	var secondScore float32
	maxTimes := 10
	deck := newDeck()
	deck.shuffle()
	for i := 0; i < maxTimes; i++ {
		card := continueGame()
		handCards, remainingDeck := deal(deck, i+1)
		if card == 'c' {
			fmt.Println("Hai preso una carta, è", handCards[i])
			cardValue := evalCard(handCards[i])
			score += cardValue
			if score > 7.5 {
				fmt.Println("Hai perso, hai superato 7 e mezzo!")
				return
			} else if score == 7.5 {
				fmt.Println("Hai vinto bravo")
				return
			} else {
				fmt.Println("Il tuo punteggio è di", score)
				fmt.Println("---------------------")
			}
		} else {
			fmt.Println("Hai fermato il gioco")
			fmt.Println("Il tuo punteggio finale è di", score)
			fmt.Println("Turno del Banco")
			bancoGame(secondScore, remainingDeck, score, i)
			return
		}
	}
}

func continueGame() rune {
	fmt.Println("Vuoi una carta? c per carta f per fermarti")
	card, err := readInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return card
}

func bancoGame(secondScore float32, deck Deck, score float32, i int) {
	maxTimes := 4
	if i == maxTimes {
		fmt.Println("Il punteggio del Banco è di", secondScore)
		fmt.Println("---------------------")
		if secondScore > score {
			fmt.Println("Il Banco ha vinto", secondScore)
			return
		}
		return
	}
	handCards, remainingDeck := deal(deck, i+1)
	fmt.Println("Il Banco ha preso una carta, è", handCards[i])
	cardValue := evalCard(handCards[i])
	secondScore += cardValue
	if secondScore > 7.5 {
		fmt.Println("Il Banco ha perso, ha sballato!, HAI VINTO!!")
		return
	} else if secondScore == 7.5 {
		fmt.Println("Il banco ha vinto", secondScore)
		return
	} else {
		fmt.Println("Il punteggio del Banco è di", secondScore)
		fmt.Println("---------------------")
		if secondScore > score {
			fmt.Println("Il Banco ha vinto", secondScore)
			return
		}
		bancoGame(secondScore, remainingDeck, score, i+1)
	}
}
