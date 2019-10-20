package enums

type Player int

const (
	Player1 Player = iota
	Player2
)

func PlayerPossibleValues() []Player {
	return []Player{
		Player1,
		Player2,
	}
}