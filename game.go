package main

import (
	"strings"
)

type GameState struct {
	PlayerName string `json:"playerName"`
	Score      int    `json:"score"`
	Lives      int    `json:"lives"`
	Hints      int    `json:"hints"`
	Level      int    `json:"level"`
	GameOver   bool   `json:"gameOver"`
	Finished   bool   `json:"finished"`
}

var game = GameState{}

func startNewGame(name string) {

	if strings.TrimSpace(name) == "" {
		name = "Player"
	}

	game = GameState{
		PlayerName: name,
		Score:      0,
		Lives:      3,
		Hints:      3,
		Level:      1,
		GameOver:   false,
		Finished:   false,
	}
}

func checkPlayerAnswer(answer string) bool {

	if game.GameOver || game.Finished {
		return false
	}

	levelIndex := game.Level - 1

	if levelIndex < 0 || levelIndex >= len(levels) {
		return false
	}

	level := levels[levelIndex]

	if normalizeAnswer(answer) == normalizeAnswer(level.Answer) {

		game.Score += level.Points

		if game.Level == len(levels) {
			game.Finished = true
		} else {
			game.Level++
		}

		return true
	}

	game.Lives--

	if game.Lives <= 0 {
		game.GameOver = true
	}

	return false
}

func getCurrentHint() string {

	if game.GameOver || game.Finished {
		return ""
	}

	if game.Hints <= 0 {
		return ""
	}

	if game.Level < 1 || game.Level > len(levels) {
		return ""
	}

	game.Hints--

	return levels[game.Level-1].Hint
}

func getCurrentLevel() Level {

	if game.Level < 1 || game.Level > len(levels) {
		return Level{}
	}

	return levels[game.Level-1]
}

func normalizeAnswer(answer string) string {

	answer = strings.ToLower(strings.TrimSpace(answer))

	answer = strings.Trim(answer, ".,!?")

	return answer
}
