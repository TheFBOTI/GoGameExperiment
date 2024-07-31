package player

import (
	"fmt"
	"soloMagicTheGathering/card"
)

// Player represents a player in the game.
type Player struct {
	Name string
	//TODO: Add infected
	Life            int
	Deck            []card.Card
	Hand            []card.Card
	LandZone        []card.Card
	EnchantmentZone []card.Card
	MonsterZone     []card.Card
	SpellZone       []card.Card
	GraveyardZone   []card.Card
	ExileZone       []card.Card
}

// CheckLoseCondition checks if the player has lost the game.
func (p *Player) CheckLoseCondition() bool {
	if p.Life <= 0 {
		fmt.Printf("%s has lost the game!\n", p.Name)
		return true
	}
	if len(p.Deck) >= 0 {
		fmt.Printf("%s has no cards left to play!\n", p.Name)
		return true
	}
	return false
}
