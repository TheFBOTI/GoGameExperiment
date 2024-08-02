package player

import (
	"fmt"
	"soloMagicTheGathering/card"
	"testing"
)

func TestLoseCondition(t *testing.T) {

	testCard := card.Card{
		Name:        "Test Card",
		Type:        "Creature",
		Description: "Test Description",
		Power:       2,
		Toughness:   3,
		ManaCost:    4,
	}

	player1 := Player{
		Name:            "Test Player",
		Life:            20,
		Deck:            []card.Card{testCard},
		Hand:            []card.Card{},
		LandZone:        []card.Card{},
		EnchantmentZone: []card.Card{},
		MonsterZone:     []card.Card{},
		SpellZone:       []card.Card{},
		GraveyardZone:   []card.Card{},
		ExileZone:       []card.Card{},
	}

	if player1.CheckLoseCondition() {
		t.Errorf("Expected false, got true for player with life > 0 and cards in deck")
	} else {
		fmt.Println("Test passed: Player with life > 0 and cards in deck has not lost.")
	}

	player2 := Player{
		Name: "Player 2",
		Life: 10,
		Deck: []card.Card{},
		Hand: []card.Card{testCard},
	}
	if !player2.CheckLoseCondition() {
		t.Errorf("Expected true, got false for player with life > 0 but no cards left")
	} else {
		fmt.Println("Test passed: Player with life > 0 but no cards left has lost.")
	}

	// Test case: Player has life <= 0
	player3 := Player{
		Name: "Player 3",
		Life: 0,
		Deck: []card.Card{{Name: "Island", Type: "Land"}},
		Hand: []card.Card{{Name: "Swamp", Type: "Land"}},
	}
	if !player3.CheckLoseCondition() {
		t.Errorf("Expected true, got false for player with life <= 0")
	} else {
		fmt.Println("Test passed: Player with life <= 0 has lost.")
	}

	player4 := Player{
		Name: "Player 4",
		Life: -1,
		Deck: []card.Card{{Name: "Plains", Type: "Land"}},
		Hand: []card.Card{{Name: "Swamp", Type: "Land"}},
	}
	if !player4.CheckLoseCondition() {
		t.Errorf("Expected true, got false for player with negative life")
	} else {
		fmt.Println("Test passed: Player with negative life has lost.")
	}
}
