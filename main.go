package main

import "fmt"

func main() {
	board := CreateBoard()
	x, y := 0, 0
	player := X
	for !board.isGameOver() {
		player = player.Other()
		pos := Position{x, y}
		board.play(player, pos)
		y++
		if y == 3 {
			x++
			y = 0
		}
		fmt.Println(board)
	}

	fmt.Println(board)
}
