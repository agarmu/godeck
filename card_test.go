package godeck

import "fmt"

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
