package godeck

import "sort"

func New(opts ...func([]Card) []Card) []Card {
	cards := make([]Card, 0)
	for _, s := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{suit: s, rank: Rank(r)})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func absRank(c Card) int {
	return int(c.suit)*int(maxRank) + int(c.rank)
}

func less(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, less(cards))
	return cards
}
