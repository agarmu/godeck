//go:generate stringer -type=Suit,Rank

package godeck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	suit Suit
	rank Rank
}

func (c Card) String() string {

	if c.suit == Joker {
		return c.suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.rank.String(), c.suit.String())

}
