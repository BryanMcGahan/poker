package game

import "math/rand"

type Deck struct {
	Cards []Card
}

func (d Deck) Initialize() Deck {
	for _, suit := range SUITS {
		for _, face := range FACES {
			card := Card{
				Value: face,
				Suit:  suit,
			}
			d.Cards = append(d.Cards, card)
		}
	}

	return d
}

func (d Deck) Shuffle() Deck {
	rand.Shuffle(len(d.Cards), func(i, j int) {d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]})
	return d
}
