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
