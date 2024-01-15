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
	ACE   cardValue = "ace"
)

var CardValueMap map[string]int = map[string]int{"two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10, "jack": 11, "queen": 12, "king": 13, "ace": 14}

var FACES [13]cardValue = [13]cardValue{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
var SUITS [4]suit = [4]suit{HEART, SPADE, CLUB, DIAMOND}

type Card struct {
	Value cardValue
	Suit  suit
}
