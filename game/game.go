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

// Play takes a channel of cards and plays the Match! game - if any two consecutive cards are matched by `Matcher`, a winner is selected according to `Judge` and the card is sent into the channel of the winning Player
func (g MatchGame) Play(cardDealer <-chan models.Card) (<-chan models.Card, <-chan models.Card) {

	p1Cards := make(chan models.Card)
	p2Cards := make(chan models.Card)

	closeChannels := func() {
		close(p1Cards)
		close(p2Cards)
	}

	// flushCards sends the supplied cards to the channel of the player
	flushCards := func(cards models.Cards, player enums.Player) {
		for _, card := range cards {
			if player == enums.Player1 {
				p1Cards <- card
			} else {
				p2Cards <- card
			}
		}
	}

	// read all dealed cards, store all unmatched cards in a buffer, and flush the cards to player's channel when there is a match
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