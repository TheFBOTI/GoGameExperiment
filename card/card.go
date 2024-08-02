package card

// Card represents a single card in the game.
type Card struct {
	Name        string
	Type        string
	Description string
	Power       int
	Toughness   int
	// Mana Cost to be improved to include different colours at a later date.
	ManaCost int
}
