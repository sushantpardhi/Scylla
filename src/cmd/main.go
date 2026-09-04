package main

import "Scylla/src/pkg/board"

func main() {
	myGame := board.NewGame()
	board.Print(myGame.WhitePieces() | myGame.BlackPieces())
}
