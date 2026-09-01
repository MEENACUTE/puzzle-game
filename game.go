package main

import "strings"

type GameState struct {
	PlayerName  string `json:"playerName"`
	Score       int    `json:"score"`
	Lives       int    `json:"lives"`
	Hints       int    `json:"hints"`
	Level       int    `json:"level"`
	Streak      int    `json:"streak"`
	BestStreak  int    `json:"bestStreak"`
	Solved      int    `json:"solved"`
	TotalLevels int    `json:"totalLevels"`
	GameOver    bool   `json:"gameOver"`
	Finished    bool   `json:"finished"`
}

var game = GameState{}

func startNewGame(name string) {

	name = strings.TrimSpace(name)

	if name == "" {
		name = "Player"
	}

	game = GameState{
		PlayerName:  name,
		Score:       0,
		Lives:       3,
		Hints:       3,
		Level:       1,
		Streak:      0,
		BestStreak:  0,
		Solved:      0,
		TotalLevels: len(levels),
		GameOver:    false,
		Finished:    false,
	}
}

func getCurrentLevel() Level {

	index := game.Level - 1

	if index < 0 || index >= len(levels) {
		return Level{}
	}

	return levels[index]
}

func checkPlayerAnswer(answer string) bool {

	if game.GameOver || game.Finished {
		return false
	}

	level := getCurrentLevel()

	if normalizeAnswer(answer) == normalizeAnswer(level.Answer) {

		game.Solved++

		game.Streak++

		if game.Streak > game.BestStreak {
			game.BestStreak = game.Streak
		}

		points := calculatePoints(level)

		game.Score += points

		if game.Level >= len(levels) {

			game.Finished = true

		} else {

			game.Level++
		}

		return true
	}

	game.Lives--

	game.Streak = 0

	if game.Lives <= 0 {
		game.GameOver = true
	}

	return false
}

func calculatePoints(level Level) int {

	points := level.Points

	// Streak bonus
	if game.Streak >= 2 {
		points += game.Streak * 10
	}

	return points
}

func getCurrentHint() string {

	if game.GameOver || game.Finished {
		return ""
	}

	if game.Hints <= 0 {
		return ""
	}

	level := getCurrentLevel()

	game.Hints--

	return level.Hint
}

func normalizeAnswer(answer string) string {

	answer = strings.ToLower(
		strings.TrimSpace(answer),
	)

	answer = strings.Trim(
		answer,
		".,!? ",
	)

	return answer
}
