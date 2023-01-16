package deck

import "fmt"

type Card struct {
	number string
	suit   Suit
}

func (card Card) Print() {
	fmt.Printf("%s %s\n", card.number, card.suit)
}
