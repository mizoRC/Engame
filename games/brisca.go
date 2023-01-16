package games

import (
	"github.com/mizoRC/engame/deck"
)

func NewBriscaGame() *deck.Deck {
	deck := deck.NewBriscaDeck()
	deck.Print()
	//log.Printf("Total number of cards: %d", len(deck.Cards))
	//log.Println("----------------------------------------------")
	deck.Shuffle()
	//deck.Print()
	//log.Println("----------------------------------------------")
	deck.Cut()
	//deck.Print()
	//log.Printf("Total number of cards: %d", len(deck.Cards))

	return deck
}
