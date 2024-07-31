package player

import (
	"soloMagicTheGathering/card"
)

// Player represents a player in the game.
type Player struct {
	Name            string
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
