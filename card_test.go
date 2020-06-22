package godeck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{rank: Ace, suit: Heart})
	fmt.Println(Card{rank: Two, suit: Club})
	fmt.Println(Card{rank: Seven, suit: Diamond})
	fmt.Println(Card{rank: Nine, suit: Club})
	fmt.Println(Card{suit: Joker})
	fmt.Println(Card{rank: Four, suit: Spade})

	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Seven of Diamonds
	// Nine of Clubs
	// Joker
	// Four of Spades
}

func TestNew(t *testing.T) {
	cards := New()

	if len(cards) != 13*4 {
		t.Error("Listlength")
	}
}
func TestSort(t *testing.T) {
	cards := New()
	cardsB := New(DefaultSort)
	if len(cards) != len(cardsB) {
		t.Error("Listlength")
	}
	for i, c := range cards {
		if c.rank != cardsB[i].rank || c.suit != cardsB[i].suit {
			t.Error("Notmatch")
		}
	}
}
