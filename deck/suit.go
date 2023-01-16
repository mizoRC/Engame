package deck

type Sweet uint
type Suit string

const (
	Oros    Suit = "Oros"
	Copas   Suit = "Copas"
	Bastos  Suit = "Bastos"
	Espadas Suit = "Espadas"
)

const (
	Hearts   Suit = "Hearts"
	Diamonds Suit = "Diamonds"
	Clubs    Suit = "Clubs"
	Spades   Suit = "Spades"
)

var BriscaSuits = []Suit{Oros, Copas, Bastos, Espadas}
var BriscaNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "Sota", "Caballo", "Rey"}

var PokerSuits = []Suit{Hearts, Diamonds, Clubs, Spades}
var PokerNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
