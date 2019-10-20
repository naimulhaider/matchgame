package enums

type MatchingCondition int

const (
	MatchingConditionFaceValue MatchingCondition = iota
	MatchingConditionSuit
	MatchingConditionFaceValueAndSuit
)

func MatchingConditionPossibleValues() []MatchingCondition {
	return []MatchingCondition{
		MatchingConditionFaceValue,
		MatchingConditionSuit,
		MatchingConditionFaceValueAndSuit,
	}
}

func MatchingConditionIsValid(mc int) bool {
	pvs := MatchingConditionPossibleValues()
	for _, pv := range pvs {
		if int(pv) == mc {
			return true
		}
	}
	return false
}