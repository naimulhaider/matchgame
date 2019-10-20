package game

import (
	"github.com/naimulhaider/matchgame/enums"
	"github.com/naimulhaider/matchgame/models"
	"math/rand"
	"time"
)

type Judge func(models.Cards) enums.Player

// EvenOddJudge returns Player1 if cards count is even, PLayer2 if cards count is odd
func EvenOddJudge(cards models.Cards) enums.Player {
	if len(cards) % 2 == 0 {
		return enums.Player1
	}
	return enums.Player2
}

// RandomJudge selects a player randomly ignoring the input parameter
func RandomJudge(cards models.Cards) enums.Player {
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	player := rander.Intn(2)
	return enums.Player(player)
}