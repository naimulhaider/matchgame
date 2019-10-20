package game

import (
	"github.com/naimulhaider/matchgame/dealer"
	"github.com/naimulhaider/matchgame/utils"
	"sync"
	"testing"
)

func TestMatchGame_Play(t *testing.T) {

	data := []struct {
		Decks int
		MaxTotal int
	} {
		{
			Decks: 1,
			MaxTotal: 1 * 52,
		},
		{
			Decks: 50,
			MaxTotal: 50 * 52,
		},
		{
			Decks: 93,
			MaxTotal: 93 * 52,
		},
		{
			Decks: 2003,
			MaxTotal: 2003 * 52,
		},
		{
			Decks: 9876,
			MaxTotal: 9876 * 52,
		},
		{
			Decks: 10001,
			MaxTotal: 10001 * 52,
		},
	}

	for _, d := range data {
		decks := utils.GenerateStandardDecks(d.Decks)
		cardDealer := dealer.ServeRandomCards(decks)

		matchGame := NewMatchGame(UseJudge(RandomJudge), UseMatcher(FaceValueAndSuitMatcher))

		p1Cards, p2Cards := matchGame.Play(cardDealer)

		p1CardsCount, p2CardsCount := 0, 0

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

		if p1CardsCount + p2CardsCount > d.MaxTotal {
			t.Errorf("Received more cards than sent, expected max: %d, got: %d", d.MaxTotal, p1CardsCount + p2CardsCount)
		}
	}
}