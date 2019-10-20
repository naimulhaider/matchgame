package utils

import (
	"fmt"
	"github.com/naimulhaider/matchgame/enums"
	"github.com/naimulhaider/matchgame/models"
	"sync"
)

// GenerateStandardDecks takes a number `n` as input and generates n standard decks
func GenerateStandardDecks(n int) (d models.Decks) {
	for i := 0; i < n; i++ {
		d  = append(d, NewStandardDeck())
	}

	return
}

// NewStandardDeck creates a new deck with all the default cards
func NewStandardDeck() (d models.Deck) {

	possibleCardValues := enums.CardValuePossibleValues()
	possibleCardSuits := enums.CardSuitPossibleValues()

	for _, val := range possibleCardValues {
		for _, suit := range possibleCardSuits {
			d.Cards = append(d.Cards, models.Card{
				Value: val,
				Suit:  suit,
			})
		}
	}

	return
}

// TallyResults counts the number of cards from the channels and declares a winner
func TallyResults(p1Cards, p2Cards <-chan models.Card) {

	var p1CardsCount, p2CardsCount int64

	w := sync.WaitGroup{}

	w.Add(2)

	go func() {
		defer w.Done()

		for _ = range p1Cards {
			p1CardsCount++
		}
	}()

	go func() {
		defer w.Done()

		for _ = range p2Cards {
			p2CardsCount++
		}
	}()

	w.Wait()

	if p1CardsCount > p2CardsCount {
		fmt.Printf("Player 1 won with %d cards\n", p1CardsCount)
	} else if p1CardsCount < p2CardsCount {
		fmt.Printf("Player 2 won with %d cards\n", p2CardsCount)
	} else {
		fmt.Printf("The game was drawn with each player having %d cards\n", p1CardsCount)
	}
}