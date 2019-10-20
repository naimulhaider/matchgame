package dealer

import (
	"github.com/naimulhaider/matchgame/utils"
	"sync"
	"testing"
)

func TestServeRandomCards(t *testing.T) {
	data := []struct {
		Decks int
		Total int
	} {
		{
			Decks: 1,
			Total: 1 * 52,
		},
		{
			Decks: 50,
			Total: 50 * 52,
		},
		{
			Decks: 93,
			Total: 93 * 52,
		},
		{
			Decks: 2003,
			Total: 2003 * 52,
		},
		{
			Decks: 9876,
			Total: 9876 * 52,
		},
		{
			Decks: 10001,
			Total: 10001 * 52,
		},
	}

	for _, d := range data {
		decks := utils.GenerateStandardDecks(d.Decks)
		cardDealer := ServeRandomCards(decks)

		cardsCount := 0

		w := sync.WaitGroup{}
		w.Add(1)

		go func() {
			defer w.Done()

			for _ = range cardDealer {
				cardsCount++
			}
		}()

		w.Wait()

		if cardsCount != d.Total {
			t.Errorf("Expected number of cards: %d, got: %d", d.Total, cardsCount)
		}
	}
}
