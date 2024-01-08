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
}

func (g Game) Setup() Game {
	g = g.setupDeck()
	return g
}

func (g Game) Start() {
	for g.Winner == (Player{}) {
		var currentHand Hand = Hand{}
		currentHand = currentHand.Setup(0, g.Players, 2, g.Deck)
		currentHand = currentHand.PlayHand()
		fmt.Println("HAND WINNER: ", currentHand.Winner.Name)
		g.Hands = append(g.Hands, currentHand)
	}
}

func (g Game) AddPlayer(newPlayer Player) Game {
	g.Players = append(g.Players, newPlayer)
	return g
}

func (g Game) setupDeck() Game {
	var deck Deck = Deck{}
	deck = deck.Initialize()

	g.Deck = deck

	return g
}
