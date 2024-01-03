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
	ONE   cardValue = "one"
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

type Card struct {
	value cardValue
	suit  suit
}
