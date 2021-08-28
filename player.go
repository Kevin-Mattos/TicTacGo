package main

type Player byte

const (
	NoPlayer Player = iota
	X
	O
	BOT
)

func (p Player) String() string {
	switch p {
	case X:
		return "X"
	case O:
		return "O"
	case BOT:
		return "B"
	default:
		return "-"
	}
}

func (player Player) Other(isBot bool) Player {
	if player == X {
		if isBot {
			return BOT
		}
		return O
	}
	return X
}
