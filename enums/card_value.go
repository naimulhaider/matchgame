package enums

type CardValue int

const (
	CardValueA CardValue = iota
	CardValue2
	CardValue3
	CardValue4
	CardValue5
	CardValue6
	CardValue7
	CardValue8
	CardValue9
	CardValue10
	CardValueJ
	CardValueQ
	CardValueK
)

func CardValuePossibleValues() []CardValue {
	return []CardValue{
		CardValueA,
		CardValue2,
		CardValue3,
		CardValue4,
		CardValue5,
		CardValue6,
		CardValue7,
		CardValue8,
		CardValue9,
		CardValue10,
		CardValueJ,
		CardValueQ,
		CardValueK,
	}
}