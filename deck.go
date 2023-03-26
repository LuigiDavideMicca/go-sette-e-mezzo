package main

import (
	"math/rand"
	"strings"
	"time"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Cuori", "Quadri", "Fiori", "Picche"}
	cardValues := []string{"Asso", "Due", "Tre", "Quattro", "Cinque", "Sei", "Sette", "Fante", "Donna", "Re"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" di "+suit)
		}
	}
	return cards
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func evalCard(c string) float32 {
	if strings.Contains(c, "Asso") {
		return 1
	} else if strings.Contains(c, "Due") {
		return 2
	} else if strings.Contains(c, "Tre") {
		return 3
	} else if strings.Contains(c, "Quattro") {
		return 4
	} else if strings.Contains(c, "Cinque") {
		return 5
	} else if strings.Contains(c, "Sei") {
		return 6
	} else if strings.Contains(c, "Sette") {
		return 7
	} else if strings.Contains(c, "Fante") {
		return 0.5
	} else if strings.Contains(c, "Donna") {
		return 0.5
	} else if strings.Contains(c, "Re di Cuori") {
		return 7
	} else {
		return 0.5
	}
}
