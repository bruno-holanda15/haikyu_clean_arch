package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"haikyu_game/internal/domain/usecase"
	"haikyu_game/internal/infra/database"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := database.NewPlayerRepository(db)
	createPlayerUsecase := usecase.NewCreatePlayerUseCase(repository)

	http.HandleFunc("/createPlayer", func(w http.ResponseWriter, r *http.Request) {
		var input usecase.InputCreatePlayer
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		output, err := createPlayerUsecase.Execute(input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(output)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running")
	http.ListenAndServe(":8081", nil)
}
