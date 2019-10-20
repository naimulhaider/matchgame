package main

import (
	"github.com/naimulhaider/matchgame/config"
	"github.com/naimulhaider/matchgame/dealer"
	"github.com/naimulhaider/matchgame/game"
	"github.com/naimulhaider/matchgame/utils"
	"log"
)

func main() {

	err := config.AskAndParseInput()
	if err != nil {
		log.Fatalf("Input Error: %s", err.Error())
	}

	decks := utils.GenerateStandardDecks(config.NoOfDecks)

	cardDealer := dealer.ServeRandomCards(decks)

	judge := game.RandomJudge
	matcher := game.MatcherFromMatchingCondition(config.MatchingCondition)
	matchGame := game.NewMatchGame(game.UseJudge(judge), game.UseMatcher(matcher))

	p1Cards, p2Cards := matchGame.Play(cardDealer)

	utils.TallyResults(p1Cards, p2Cards)
}

