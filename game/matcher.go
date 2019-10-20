package game

import (
	"github.com/naimulhaider/matchgame/enums"
	"github.com/naimulhaider/matchgame/models"
)

type Matcher func(models.Card, models.Card) bool

func FaceValueMatcher(a models.Card, b models.Card) bool {
	return a.Value == b.Value
}

func SuitMatcher(a models.Card, b models.Card) bool {
	return a.Suit == b.Suit
}

func FaceValueAndSuitMatcher(a models.Card, b models.Card) bool {
	return a.Value == b.Value && a.Suit == b.Suit
}

func MatcherFromMatchingCondition(matchingCondition enums.MatchingCondition) Matcher {
	switch matchingCondition {
	case enums.MatchingConditionFaceValue:
		return FaceValueMatcher
	case enums.MatchingConditionSuit:
		return SuitMatcher
	case enums.MatchingConditionFaceValueAndSuit:
		return FaceValueAndSuitMatcher
	default:
		return FaceValueAndSuitMatcher
	}
}