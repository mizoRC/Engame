package deck

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewBriscaDeck() *Deck {
	deck := new(Deck)
	for _, suit := range BriscaSuits {
		for _, number := range BriscaNumbers {
			card := Card{number: number, suit: suit}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (deck *Deck) Print() {
	for _, card := range deck.Cards {
		card.Print()
	}
}

func (deck *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})
}

func (deck *Deck) Cut() {
	var halfs []Card
	chunkSize := len(deck.Cards) / 2

	firstHalf := deck.Cards[0:chunkSize]
	secondHalf := deck.Cards[chunkSize:len(deck.Cards)]

	halfs = append(halfs, secondHalf...)
	halfs = append(halfs, firstHalf...)
	deck.Cards = halfs
}
