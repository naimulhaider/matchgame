package config

import (
	"fmt"
	"github.com/naimulhaider/matchgame/enums"
)

var (
	NoOfDecks int
	MatchingCondition enums.MatchingCondition
)

func AskAndParseInput() error {
	fmt.Println("Please enter the number of decks to use: ")
	n, err := fmt.Scanln(&NoOfDecks)
	if err != nil {
		return fmt.Errorf("failed to read number of decks: %s", err.Error())
	}
	if n > 1 {
		return fmt.Errorf("failed to read number of decks, expected one input value, got: %d", n)
	}

	var matchingCondition int
	fmt.Println("Please select the Matching Condition\n(1) Face Value of Card\n(2) Suit of the Card\n(3) Both Face Value and Suit\nEnter your selection: ")
	n, err = fmt.Scanln(&matchingCondition)
	if err != nil {
		return fmt.Errorf("failed to read matching condition: %s", err.Error())
	}
	if n > 1 {
		return fmt.Errorf("failed to read matching condition, expected one input value, got: %d", n)
	}

	// check if input matching condition was a valid number
	isValid := enums.MatchingConditionIsValid(matchingCondition-1)
	if !isValid {
		return fmt.Errorf("the entered value %d is not a valid matching condition number", matchingCondition)
	}

	MatchingCondition = enums.MatchingCondition(matchingCondition-1)

	return nil
}