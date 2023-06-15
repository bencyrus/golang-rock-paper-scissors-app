package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	if err := renderTemplate(w, "index.html"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, err := strconv.Atoi(r.URL.Query().Get("c"))
	if err != nil {
		http.Error(w, "Invalid choice", http.StatusBadRequest)
		return
	}
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	log.Println("Server started on: http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func renderTemplate(w http.ResponseWriter, page string) error {
	t, err := template.ParseFiles(page)
	if err != nil {
		return err
	}

	err = t.Execute(w, nil)
	if err != nil {
		return err
	}

	return nil
}
