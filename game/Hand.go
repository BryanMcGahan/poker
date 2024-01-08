package game

import "fmt"

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
	CommunityCards PlayedCards
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
)

// order of a hand
// 1. Deal
// 2. Bid
// 3. Flop
// 4. Bid
// 5. Turn
// 6. Bid
// 7. River
// 8. Bid
// 9. Determine winner

func (h Hand) PlayHand() Hand {
	h = h.Deal()
	for h.Winner == (Player{}) {
		h = h.BidRound()
		h = h.CardFlip()

		if h.Stage == DETER {
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
	h.LittleBlind = h.Dealer + 1
	h.BigBlind = h.LittleBlind + 1
	h.MinBid = minBid
	h.Deck = deck
	h.Stage = FLOP

	fmt.Println(h.CommunityCards)
	for i := 0; i < 7; i++ {
		h.Deck = h.Deck.Shuffle()
	}

	return h
}

func (h Hand) Deal() Hand {
	for i := 0; i < 2; i++ {
		j := h.Dealer + 1
		sameRound := true
		for sameRound {
			if h.Players[j].hand.card1 == (Card{}) {
				h.Players[j].hand.card1 = h.Deck.Cards[0]
				h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[0+1:]...)
			} else {
				h.Players[j].hand.card2 = h.Deck.Cards[0]
				h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[0+1:]...)
			}

			if j == len(h.Players)-1 {
				j = 0
			} else if j == h.Dealer {
				sameRound = false
			} else {
				j++
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
		h.CommunityCards.card1 = h.Deck.Cards[0]
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		h.CommunityCards.card2 = h.Deck.Cards[0]
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		h.CommunityCards.card3 = h.Deck.Cards[0]
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		break
	case TURN:
		h.Stage = RIVER
		h.CommunityCards.card4 = h.Deck.Cards[0]
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
		fmt.Println(h.Stage)
		break
	case RIVER:
		h.Stage = DETER
		h.CommunityCards.card5 = h.Deck.Cards[0]
		h.Deck.Cards = append(h.Deck.Cards[:0], h.Deck.Cards[1:]...)
	}

	return h
}

func (h Hand) BidRound() Hand {
	fmt.Println("BID")
	var i int
	var exit int
	var allPlayersBid bool = false
	switch h.Stage {
	case FLOP:
		i = h.BigBlind + 1
		exit = h.BigBlind
		break
	default:
		i = h.Dealer + 1
		exit = h.Dealer
		break
	}
	for !allPlayersBid {

		player, bidAmount := h.Players[i].PlayBid(h.CurrentBid)

		h.CurrentBid = bidAmount
		h.Players[i] = player

		switch h.Players[i].CurrentPlay {
		case CALL, BID, RAISE:
			h.Players[i] = h.Players[i].Bid(bidAmount)
			h.Pot += bidAmount
			break
		case FOLD:
			fmt.Println(h.Players)
			h.Players = append(h.Players[:i], h.Players[i+1:]...)
			i--
			fmt.Println(h.Players)
			break
		}

		if i == exit {
			allPlayersBid = true
		} else if i == len(h.Players)-1 {
			i = 0
		} else {
			i++
		}
	}

	h.CurrentBid = 0

	return h
}
