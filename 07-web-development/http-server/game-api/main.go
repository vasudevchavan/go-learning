package main

import (
	"encoding/json"
	"fmt"
	"game/game"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemple(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := game.PlayRound(playerChoice)
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {

	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	fmt.Println("Starting web server on port 8081")
	http.ListenAndServe(":8081", nil)
}

func renderTemple(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
