package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func main() {

	http.Handle(
		"/",
		http.FileServer(http.Dir("./web")),
	)

	http.HandleFunc("/api/start", startGameHandler)

	http.HandleFunc("/api/state", stateHandler)

	http.HandleFunc("/api/answer", answerHandler)

	http.HandleFunc("/api/hint", hintHandler)

	http.HandleFunc("/api/level", levelHandler)

	fmt.Println("======================================")
	fmt.Println("       THE THREE PUZZLES")
	fmt.Println("======================================")
	fmt.Println()
	fmt.Println("Server running at:")
	fmt.Println("http://localhost:8080")
	fmt.Println()
	fmt.Println("Opening browser...")

	go openBrowser("http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func startGameHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var request struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			"Invalid request",
			http.StatusBadRequest,
		)

		return
	}

	startNewGame(request.Name)

	sendJSON(w, game)
}

func stateHandler(w http.ResponseWriter, r *http.Request) {

	sendJSON(w, game)
}

func answerHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	var request struct {
		Answer string `json:"answer"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(
			w,
			"Invalid request",
			http.StatusBadRequest,
		)

		return
	}

	correct := checkPlayerAnswer(request.Answer)

	response := struct {
		Correct bool      `json:"correct"`
		Game    GameState `json:"game"`
	}{
		Correct: correct,
		Game:    game,
	}

	sendJSON(w, response)
}

func hintHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(
			w,
			"Method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	hint := getCurrentHint()

	response := struct {
		Hint string    `json:"hint"`
		Game GameState `json:"game"`
	}{
		Hint: hint,
		Game: game,
	}

	sendJSON(w, response)
}

func sendJSON(w http.ResponseWriter, data interface{}) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(data)
}

func openBrowser(url string) {

	var command *exec.Cmd

	switch runtime.GOOS {

	case "windows":

		command = exec.Command(
			"cmd",
			"/c",
			"start",
			"",
			url,
		)

	case "darwin":

		command = exec.Command(
			"open",
			url,
		)

	default:

		command = exec.Command(
			"xdg-open",
			url,
		)
	}

	command.Run()
}

func levelHandler(w http.ResponseWriter, r *http.Request) {

	level := getCurrentLevel()

	response := struct {
		Number      int    `json:"number"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Question    string `json:"question"`
		Points      int    `json:"points"`
	}{
		Number:      level.Number,
		Title:       level.Title,
		Description: level.Description,
		Question:    level.Question,
		Points:      level.Points,
	}

	sendJSON(w, response)
}
