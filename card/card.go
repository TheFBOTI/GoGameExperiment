package card

// Card represents a single card in the game.
type Card struct {
	Name        string
	Type        string
	Description string
	Power       int
	Toughness   int
	ManaCost    int
}
