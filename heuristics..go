package main

import "fmt"

// type UnplayableBoard interface {
// 	getPos() Player
// 	setPos(pos Position)
// 	isValidPosition(pos Position) bool
// }

func getBestPos(node Board, player Player) Position {
	num_jog := todas_pos(&node)
	var a []int
	if len(num_jog) == 9 {
		return Position{0, 0}
	}

	for _, jog := range num_jog {
		copy := getCopy(&node)
		copy.setPos(jog, BOT)
		a = append(a, minmax(copy, 0, false, X))
	}

	var besPost Position
	var best_value = -9999

	for index := range num_jog {
		if a[index] >= best_value {
			besPost = num_jog[index]
			best_value = a[index]
		}
	}

	fmt.Println(num_jog)
	fmt.Println(a)

	return besPost
}

func minmax(node *Board, profundidade int, isMaxPlayer bool, player Player) int {
	v := 0
	best_value := 0
	aux := todas_pos(node)

	if len(aux) == 0 || node.checkVictory(player) || node.checkVictory(player.Other(true)) {
		if node.checkVictory(BOT) {
			return 10
		}
		if node.checkVictory(BOT.Other(true)) {
			return -10
		}
		return 0
	}
	if isMaxPlayer {
		best_value = -10
	} else {
		best_value = 10
	}
	for _, pos := range aux {

		copy := getCopy(node)
		copy.setPos(pos, player)
		v = minmax(copy, profundidade+1, !isMaxPlayer, player.Other(true))

		if isMaxPlayer {
			best_value = maior(best_value, v)
		} else {
			best_value = menor(best_value, v)
		}
	}

	return best_value
}

func todas_pos(board *Board) []Position {
	var a []Position
	for x := range board.Board {
		for y := range board.Board[x] {
			if NoPlayer == *board.getPos(Position{x, y}) {
				a = append(a, Position{x, y})
			}
		}
	}

	return a
}

func setPos(b Board, p Position, pl Player) Board {
	b.setPos(p, pl)
	return b
}

func maior(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func menor(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func getCopy(b *Board) *Board {
	newBoard := Board{}
	for x := range b.Board {
		for y := range b.Board[x] {
			newBoard.setPos(Position{x, y}, *b.getPos(Position{x, y}))
		}
	}
	return &newBoard
}
