package main

import (
	"poker/game"
)

func main() {
	var currentGame game.Game = game.Game{}
	currentGame = currentGame.Setup()

	var player1 game.Player = game.Player{
		Name: "Bryan McGahan",
	}

	var player2 game.Player = game.Player{
		Name: "Easton Clark",
	}

	var player3 game.Player = game.Player{
		Name: "Josiah Connell",
	}

	var player4 game.Player = game.Player{
		Name: "Brady Roberts",
	}

	currentGame = currentGame.AddPlayer(player1)
	currentGame = currentGame.AddPlayer(player2)
	currentGame = currentGame.AddPlayer(player3)
	currentGame = currentGame.AddPlayer(player4)

	currentGame.Start()
}
