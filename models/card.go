package models

import "github.com/naimulhaider/matchgame/enums"

type Cards []Card

type Card struct {
	Value enums.CardValue
	Suit enums.CardSuit
}
