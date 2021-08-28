package main

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
	if !pos.validate() || !b.isValidPosition(pos) {
		return
	}
	b.setPos(pos, p)
	b.TimesPlayed++
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

func (b *Board) isGameOver() bool {
	if b.TimesPlayed == 9 {
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
