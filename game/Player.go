package game

import (
	"fmt"
)

type Player struct {
	Name        string
	hand        hand
	Chips       int
	CurrentPlay string
	Folded      bool
	// NextPlayer *Player
}

type hand struct {
	card1 Card
	card2 Card
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
