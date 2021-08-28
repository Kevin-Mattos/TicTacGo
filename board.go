package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Board       [3][3]Player
	TimesPlayed int
}

func (b Board) String() string {
	var a string
	for _, line := range b.Board {
		for _, player := range line {
			a += player.String()
		}
		a += "\n"
	}
	return a

}

func (b *Board) play(p Player, pos Position) {

	b.setPos(pos, p)
	b.TimesPlayed++
}

func (b *Board) valitePlay(pos Position) bool {
	if !pos.validate() || !b.isValidPosition(pos) {
		return false
	}
	return true
}

func (b *Board) isValidPosition(pos Position) bool {
	return b.Board[pos.X][pos.Y] == NoPlayer
}

func (b *Board) getPos(pos Position) *Player {
	return &b.Board[pos.X][pos.Y]
}

func (b *Board) setPos(pos Position, p Player) {
	*b.getPos(pos) = p
}

func (b *Board) isGameOver(p Player) bool {
	if b.TimesPlayed == 9 || b.checkVictory(p) {
		return true
	}

	return false
}

func CreateBoard() Board {
	return Board{
		Board: [3][3]Player{
			{
				NoPlayer, NoPlayer, NoPlayer,
			},
			{
				NoPlayer, NoPlayer, NoPlayer,
			},
			{
				NoPlayer, NoPlayer, NoPlayer,
			},
		},
	}
}

func (board Board) getPlay(player Player) Position {
	var pos Position = Position{-3, -3}
	var x, y int
	//todo validar falha do parse
	if player == X || player == O {
		for !board.valitePlay(pos) {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			x, _ = strconv.Atoi(text)
			text, _ = reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			y, _ = strconv.Atoi(text)
			pos = Position{x, y}
			if !board.valitePlay(pos) {
				fmt.Println("digite novamente")
			}
		}

		return pos
	}
	pos = getBestPos(board, BOT)
	// fmt.Println(pos)
	return pos
}

func (board *Board) checkVictory(player Player) bool {
	win := 0
	for _, line := range board.Board {
		for _, boardPlayer := range line {
			if player == boardPlayer {
				win += 1
			}
			if win == 3 {
				return true
			}
		}
		win = 0
	}

	for x := range board.Board {
		for y := range board.Board[x] {
			if player == *board.getPos(Position{y, x}) {
				win += 1
			}
			if win == 3 {
				return true
			}
		}
		win = 0
	}

	for x := range board.Board {
		if player == *board.getPos(Position{x, x}) {
			win += 1
		}
	}
	if win == 3 {
		return true
	}
	win = 0

	for x := range board.Board {
		if player == *board.getPos(Position{x, 2 - x}) {
			win += 1
		}
	}
	if win == 3 {
		return true
	}
	win = 0

	return false
}
