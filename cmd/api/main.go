package main

import (
	"database/sql"
	"fmt"
	"haikyu_game/internal/infra/database"
	"haikyu_game/internal/infra/http/echo"
	"haikyu_game/internal/infra/webserver"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	repository := database.NewPlayerRepository(db)
	server := echo.Handlers(repository)

	fmt.Println("Server running")
	webserver.Start("8081", server)
}
