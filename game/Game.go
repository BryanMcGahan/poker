package game

import "fmt"

const (
	CHECK = "CHECK"
	BID   = "BID"
	RAISE = "RAISE"
	FOLD  = "FOLD"
	CALL  = "CALL"
	NONE  = "NONE"
)

type Game struct {
	Deck        Deck
	Players     []Player
	Dealer      int // index into players array
	LittleBlind int
	BigBlind    int
	Pot         int
	Bid         int
	Round       int
	Hands       []Hand
	Winner      Player
	BuyIn       int
}

func (g Game) Setup() Game {
	g.BuyIn = 40
	return g
}

func (g Game) Start() {
	var dealer int = 0
	for g.Winner == (Player{}) {
		var currentHand Hand = Hand{}
		g = g.setupDeck()
		currentHand = currentHand.Setup(dealer, g.Players, 2, g.Deck)
		currentHand = currentHand.PlayHand()
		g.Hands = append(g.Hands, currentHand)
		dealer = (dealer + 1) % len(g.Players)
	}
}

func (g Game) AddPlayer(newPlayer Player) Game {
	newPlayer.Chips = g.BuyIn
	fmt.Println(newPlayer)
	g.Players = append(g.Players, newPlayer)
	return g
}

func (g Game) setupDeck() Game {
	var deck Deck = Deck{}
	deck = deck.Initialize()

	g.Deck = deck

	return g
}
