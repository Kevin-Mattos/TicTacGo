package main

type Player byte

const (
	NoPlayer Player = iota
	X
	O
)

func (p Player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "o"
	default:
		return "-"
	}
}

func (player Player) Other() Player {
	if player == X {
		return O
	}
	return X
}
