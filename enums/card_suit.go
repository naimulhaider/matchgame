package enums

type CardSuit int

const (
	CardSuitSpade CardSuit = iota
	CardSuitClub
	CardSuitDiamond
	CardSuitHeart
)

func CardSuitPossibleValues() []CardSuit {
	return []CardSuit{
		CardSuitSpade,
		CardSuitClub,
		CardSuitDiamond,
		CardSuitHeart,
	}
}