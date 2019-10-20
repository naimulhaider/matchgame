package game

type GameOption func(*MatchGame)

func UseJudge(judge Judge) GameOption {
	return func(game *MatchGame) {
		game.judge = judge
	}
}

func UseMatcher(matcher Matcher) GameOption {
	return func(game *MatchGame) {
		game.isMatch = matcher
	}
}