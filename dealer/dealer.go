package dealer

import (
	"github.com/naimulhaider/matchgame/models"
	"math/rand"
	"time"
)

// ServeRandomCards takes an array of decks and returns a channel in which random cards are sent
func ServeRandomCards(decks models.Decks) <-chan models.Card {

	allCards := models.Cards{}

	// accumulate all cards into one array
	for _, deck := range decks {
		for _, card := range deck.Cards {
			allCards = append(allCards, card)
		}
	}

	// shuffle the array to randomize
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	rander.Shuffle(len(allCards), func(i, j int) {
		allCards[i], allCards[j] = allCards[j], allCards[i]
	})

	cardsPipe := make(chan models.Card)

	// send the shuffled cards to the channel and close after
	go func() {
		defer close(cardsPipe)

		for _, card := range allCards {
			cardsPipe <- card
		}
	}()

	return cardsPipe
}