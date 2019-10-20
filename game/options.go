package game

type GameOption func(*MatchGame)

// UseJudge is a GameOption that sets the judge used in the Game to select the winner for each round
func UseJudge(judge Judge) GameOption {
	return func(game *MatchGame) {
		game.judge = judge
	}
}

// UseMatcher is a GameOption that sets the matcher used in the Game
func UseMatcher(matcher Matcher) GameOption {
	return func(game *MatchGame) {
		game.isMatch = matcher
	}
}