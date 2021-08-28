package main

import "fmt"

func main() {
	board := CreateBoard()
	player := X
	var pos Position
	for !board.isGameOver() {
		player = player.Other()
		pos = board.getPlay(player)
		board.play(player, pos)
		fmt.Println(board)
	}

	fmt.Println(board)
}
