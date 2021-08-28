package main

type Position struct {
	X, Y int
}

func (p *Position) validate() bool {
	min := 0
	max := 2
	return p.X >= min && p.X <= max &&
		p.Y >= min && p.Y <= max
}

func (pos *Position) isCorner() bool {
	return (pos.X == 0 && (pos.Y == 0 || pos.Y == 2)) || (pos.Y == 2 && (pos.Y == 0 || pos.Y == 2))
}

func (pos *Position) isCenter() bool {
	return pos.X == 1 && pos.Y == 1
}
