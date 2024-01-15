package game

import (
	"cmp"
	"fmt"
	"slices"
)

type Player struct {
	Name        string
	hand        hand
	Chips       int
	CurrentPlay string
	HandRank    int
	HighCard    Card
	BestFive    Plays
	// NextPlayer *Player
}

type Plays string

type hand struct {
	card1 Card
	card2 Card
}

const (
	HIGHCARD      = "HIGHCARD"
	PAIR          = "PAIR"
	TWOPAIR       = "TWOPAIR"
	THREEOFKIND   = "THREEOFKIND"
	STRAIGHT      = "STRAIGHT"
	FLUSH         = "FLUSH"
	FULLHOUSE     = "FULLHOUSE"
	FOUROFKIND    = "FOUROFKIND"
	STRAIGHTFLUSH = "STRAIGHTFLUSH"
	ROYALFLUSH    = "ROYALFLUSH"
)

func (p Player) RankHand(communityCards []Card) Player {

	if CardValueMap[string(p.hand.card1.Value)] > CardValueMap[string(p.hand.card2.Value)] {
		p.HighCard = p.hand.card1
	} else {
		p.HighCard = p.hand.card2
	}

	communityCards = append(communityCards, p.hand.card1)
	communityCards = append(communityCards, p.hand.card2)

	slices.SortFunc[[]Card](communityCards, func(card1, card2 Card) int {
		return cmp.Compare(CardValueMap[string(card1.Value)], CardValueMap[string(card2.Value)])
	})

	var continuous int = 1
	for i, card := range communityCards {
		if communityCards[i].Value == TWO && communityCards[len(communityCards)-1].Value == ACE {
			continuous++
			continue
		}
		if i < len(communityCards)-1 {
			if continuous == 5 {
				break
			}
			if CardValueMap[string(card.Value)]+1 == CardValueMap[string(communityCards[i+1].Value)] {
				continuous++
			} else {
				if card.Value == ACE && communityCards[i+1].Value == TWO {
					continuous++
				} else {
					continuous = 1
				}
			}
		}
	}

	if continuous == 5 {
		p.BestFive = STRAIGHT
	}

	var card1Pairs []Card
	var card2Pairs []Card
	var boardPairs []Card
	var boardSuitMatches []Card
	var boardSuits map[string]int = map[string]int{string(HEART): 0, string(CLUB): 0, string(SPADE): 0, string(DIAMOND): 0}

	for i, card := range communityCards {

		// Need to move this to the hand
		switch card.Suit {
		case HEART:
			boardSuits[string(HEART)]++
			break
		case CLUB:
			boardSuits[string(CLUB)]++
			break
		case SPADE:
			boardSuits[string(SPADE)]++
			break
		case DIAMOND:
			boardSuits[string(DIAMOND)]++
			break
		}

		// This doesn't need to be calculated for every player
		if i < len(communityCards)-1 {
			for j := i + 1; j < len(communityCards); j++ {
				if card.Value == communityCards[j].Value {
					boardPairs = append(boardPairs, card)
				}
				if card.Suit == communityCards[j].Suit {
					boardSuitMatches = append(boardSuitMatches, card)
				}
			}
		}

		var pairCard1 bool = card.Value == p.hand.card1.Value
		var pairCard2 bool = card.Value == p.hand.card2.Value

		if pairCard1 {
			card1Pairs = append(card1Pairs, p.hand.card1)
		}

		if pairCard2 {
			card2Pairs = append(card2Pairs, p.hand.card2)
		}

	}

	fmt.Println()
	fmt.Println(boardSuits)
	fmt.Println(boardPairs)
	fmt.Println(p.Name, p.hand)
	fmt.Println(card1Pairs)
	fmt.Println(card2Pairs)

	return p
}

func (p Player) Bid(bidAmound int) Player {
	p.Chips -= bidAmound
	return p
}

func (p Player) PlayBid(handCurrentBid int) (Player, int) {
	var bidAmount int = handCurrentBid
	fmt.Println("What would you like to do?", p.Name)
	if handCurrentBid > 0 {
		fmt.Println()
		fmt.Println("1. Call")
		fmt.Println("2. Raise")
		fmt.Println("3. Fold")
		fmt.Print("Play: ")
		var playerOption int

		fmt.Scan(&playerOption)
		switch playerOption {
		case 1:
			p.CurrentPlay = CALL
			break
		case 2:
			p.CurrentPlay = RAISE
			fmt.Println()
			fmt.Print("Amount: ")
			fmt.Scan(&bidAmount)
			break
		case 3:
			p.CurrentPlay = FOLD
			break
		}
	} else {
		fmt.Println()
		fmt.Println("1. Check")
		fmt.Println("2. Bid")
		fmt.Println("3. Fold")
		fmt.Print("Play: ")
		var playerOption int

		fmt.Scan(&playerOption)
		switch playerOption {
		case 1:
			p.CurrentPlay = CHECK
			break
		case 2:
			p.CurrentPlay = BID
			fmt.Println()
			fmt.Print("Amount: ")
			fmt.Scan(&bidAmount)
			break
		case 3:
			p.CurrentPlay = FOLD
			break
		}
	}

	return p, bidAmount
}

func (p Player) PrintPlayerInfo() {
	fmt.Println()
	fmt.Println(p.Name)
	fmt.Println(p.hand)
	fmt.Println()
}

func (p Player) dealFirstCard(card Card) Player {
	p.hand.card1 = card
	return p
}

func (p Player) dealSecondCard(card Card) Player {
	p.hand.card2 = card
	return p
}
