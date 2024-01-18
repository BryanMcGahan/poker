package game

import (
	"cmp"
	"fmt"
	"slices"
)

type Hand struct {
	Players        []Player
	Stage          string
	Pot            int
	MinBid         int
	CurrentBid     int
	Dealer         int
	LittleBlind    int
	BigBlind       int
	Deck           Deck
	CommunityCards []Card
	Winner         Player
}

type PlayedCards struct {
	card1 Card
	card2 Card
	card3 Card
	card4 Card
	card5 Card
}

const (
	SETUP = "SETUP"
	FLOP  = "FLOP"
	TURN  = "TURN"
	RIVER = "RIVER"
	DETER = "DETER"
	LAST  = "LAST"
)

// order of a hand
// 1. Deal
// 2. Bid
// 3. Flop
// 4. Bid
// 5. Turn
// 6. Bid
//$7. River
// 8. Bid
// 9. Determine winner

func (h Hand) DetermineWinner() Hand {

	var boardsCardsSorted []Card = h.CommunityCards

	slices.SortFunc[[]Card](boardsCardsSorted, func(card1, card2 Card) int {
		return cmp.Compare(CardValueMap[string(card1.Value)], CardValueMap[string(card2.Value)])
	})

	var boardPairs []Card = getBoardPairs(boardsCardsSorted)
	fmt.Println("Board Pairs", boardPairs)

	var boardSuits map[string]int = getBoardSuitCounts(boardsCardsSorted)
	fmt.Println("Board suits", boardSuits)

	return h
}

func getBoardSuitCounts(cards []Card) map[string]int {

	var suitCounts map[string]int = map[string]int{string(HEART): 0, string(DIAMOND): 0, string(SPADE): 0, string(CLUB): 0}

	for i := 0; i < len(cards); i++ {
		suitCounts[string(cards[i].Suit)]++
	}

	return suitCounts
}

func getBoardPairs(cards []Card) []Card {

	var pairs []Card
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if CardValueMap[string(cards[i].Value)] == CardValueMap[string(cards[j].Value)] {
				pairs = append(pairs, cards[i])
			}
		}
	}

	return pairs
}

func (h Hand) PlayHand() Hand {
	h = h.Deal()
	for h.Winner == (Player{}) {
		h = h.BidRound()
		h = h.CardFlip()
		if h.Stage == DETER {
			h = h.DetermineWinner()
			h.Winner = h.Players[0]
		}
	}

	return h
}

func (h Hand) showPlayersHands() {
	for i := 0; i < len(h.Players); i++ {
		h.Players[i].PrintPlayerInfo()
	}
}

func (h Hand) Setup(dealer int, activePlayers []Player, minBid int, deck Deck) Hand {

	h.Players = activePlayers
	h.Dealer = dealer
	h.LittleBlind = (h.Dealer + 1) % len(h.Players)
	h.BigBlind = (h.LittleBlind + 1) % len(h.Players)
	h.MinBid = minBid
	h.Deck.Cards = deck.Cards
	h.Stage = FLOP

	for i := 0; i < 7; i++ {
		h.Deck = h.Deck.Shuffle()
	}

	return h
}

func (d Deck) printDeck() {
	for i, card := range d.Cards {
		fmt.Println(i, card)
	}
}

func (h Hand) Deal() Hand {
	for i := 0; i < 2; i++ {
		j := h.LittleBlind
		sameRound := true
		for sameRound {
			if h.Players[j].hand.card1 == (Card{}) {
				h.Players[j].hand.card1 = h.Deck.Cards[0]
				h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[0+1:]...)
			} else {
				h.Players[j].hand.card2 = h.Deck.Cards[0]
				h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[0+1:]...)
			}

			if j == h.Dealer {
				sameRound = false
			} else {
				j = (j + 1) % len(h.Players)
			}
		}
	}
	return h
}

func (h Hand) CardFlip() Hand {
	h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
	switch h.Stage {
	case FLOP:
		h.Stage = TURN
		h.CommunityCards = append(h.CommunityCards, h.Deck.Cards[0])
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		h.CommunityCards = append(h.CommunityCards, h.Deck.Cards[0])
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		h.CommunityCards = append(h.CommunityCards, h.Deck.Cards[0])
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		break
	case TURN:
		h.Stage = RIVER
		h.CommunityCards = append(h.CommunityCards, h.Deck.Cards[0])
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		break
	case RIVER:
		h.Stage = LAST
		h.CommunityCards = append(h.CommunityCards, h.Deck.Cards[0])
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
	}

	return h
}

func (h Hand) BidRound() Hand {
	var i int
	var exit int
	var allPlayersBid bool = false
	switch h.Stage {
	case FLOP:
		i = (h.BigBlind + 1) % len(h.Players)
		exit = h.BigBlind
		break
	default:
		i = (h.Dealer + 1) % len(h.Players)
		exit = h.Dealer
		if h.Stage == LAST {
			h.Stage = DETER
		}
		break
	}

	for !allPlayersBid {

		player, bidAmount := h.Players[i].PlayBid(h.CurrentBid)

		h.CurrentBid = bidAmount
		h.Players[i] = player

		switch h.Players[i].CurrentPlay {
		case CALL, BID:
			h.Players[i] = h.Players[i].Bid(bidAmount)
			h.Pot += bidAmount
			break
		case RAISE:
			h.Players[i] = h.Players[i].Bid(bidAmount)
			h.Pot += bidAmount
			if i == 0 {
				exit = len(h.Players) - 1
			} else {
				exit = i - 1
			}
			break
		case FOLD:
			h.Players = append(h.Players[:i], h.Players[i+1:]...)
			i--
			break
		default:
			break
		}

		if i == exit {
			allPlayersBid = true
		} else {
			i = (i + 1) % len(h.Players)
		}
	}

	h.CurrentBid = 0

	return h
}
