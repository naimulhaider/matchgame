package game

import (
	"github.com/naimulhaider/matchgame/enums"
	"github.com/naimulhaider/matchgame/models"
)

type Matcher func(models.Card, models.Card) bool

// FaceValueMatcher matches two cards by Value
func FaceValueMatcher(a models.Card, b models.Card) bool {
	return a.Value == b.Value
}

// SuitMatcher matches two cards by Suit
func SuitMatcher(a models.Card, b models.Card) bool {
	return a.Suit == b.Suit
}

// FaceValueAndSuitMatcher matches two cards by both Face Value and Suit
func FaceValueAndSuitMatcher(a models.Card, b models.Card) bool {
	return a.Value == b.Value && a.Suit == b.Suit
}

// MatcherFromMatchingCondition takes a MatchingCondition and returns the corresponding Matcher function
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