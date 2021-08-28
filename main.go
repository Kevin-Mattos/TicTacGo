package main

import "fmt"

func main() {
	board := CreateBoard()
	player := X
	var pos Position
	isBot := true
	// board.setPos(Position{0, 0}, player.Other(isBot))
	// board.setPos(Position{2, 2}, BOT)
	// board.setPos(Position{1, 1}, player)
	// board.setPos(Position{1, 1}, X)
	// board.setPos(Position{1,1}, X)
	// board.setPos(Position{1,1}, X)

	fmt.Println(board)
	for !board.isGameOver(player) {
		player = player.Other(isBot)
		pos = board.getPlay(player)
		board.play(player, pos)
		fmt.Println(board)

	}

	fmt.Println(board)
}
