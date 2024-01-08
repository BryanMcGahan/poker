package game

type suit string

const (
	HEART   suit = "heart"
	SPADE   suit = "spade"
	CLUB    suit = "club"
	DIAMOND suit = "diamond"
)

type cardValue string

const (
	ACE   cardValue = "ace"
	TWO   cardValue = "two"
	THREE cardValue = "three"
	FOUR  cardValue = "four"
	FIVE  cardValue = "five"
	SIX   cardValue = "six"
	SEVEN cardValue = "seven"
	EIGHT cardValue = "eight"
	NINE  cardValue = "nine"
	TEN   cardValue = "ten"
	JACK  cardValue = "jack"
	QUEEN cardValue = "queen"
	KING  cardValue = "king"
)

var FACES [13]cardValue = [13]cardValue{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
var SUITS [4]suit = [4]suit{HEART, SPADE, CLUB, DIAMOND}

type Card struct {
	Value cardValue
	Suit  suit
}
