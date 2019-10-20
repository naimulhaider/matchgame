package game

import (
	"github.com/naimulhaider/matchgame/enums"
	"github.com/naimulhaider/matchgame/models"
)

type MatchGame struct {
	judge Judge
	isMatch Matcher
}

func NewMatchGame(opts ...GameOption) MatchGame {
	matchGame := MatchGame{}

	for _, op := range opts {
		op(&matchGame)
	}

	return matchGame
}

func (g MatchGame) Play(cardDealer <-chan models.Card) (<-chan models.Card, <-chan models.Card) {

	p1Cards := make(chan models.Card)
	p2Cards := make(chan models.Card)

	closeChannels := func() {
		close(p1Cards)
		close(p2Cards)
	}

	// flushCards sends the cards to the channel for the player
	flushCards := func(cards models.Cards, player enums.Player) {
		for _, card := range cards {
			if player == enums.Player1 {
				p1Cards <- card
			} else {
				p2Cards <- card
			}
		}
	}

	go func() {
		defer closeChannels()

		cardsBuffer := models.Cards{}

		for card := range cardDealer {
			bufferLen := len(cardsBuffer)
			if bufferLen > 0 && g.isMatch(cardsBuffer[bufferLen-1], card) {
				winner := g.judge(cardsBuffer)
				flushCards(cardsBuffer, winner)
				cardsBuffer = models.Cards{}
			} else {
				cardsBuffer = append(cardsBuffer, card)
			}
		}
	}()

	return p1Cards, p2Cards
}