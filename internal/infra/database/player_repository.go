package database

import (
	"database/sql"
	"haikyu_game/internal/domain/player"
)

type PlayerRepository struct {
	Db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *PlayerRepository {
	return &PlayerRepository{
		Db: db,
	}
}

func (p *PlayerRepository) Save(player *player.Player) error {
	stmt, err := p.Db.Prepare("INSERT INTO players (id, name, position, height, play_grade) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(player.ID, player.Name, player.Position, player.Height, player.PlayGrade)
	if err != nil {
		return err
	}

	return nil
}

func (p *PlayerRepository) ListAll() ([]player.Player, error) {
	rows, err := p.Db.Query("SELECT id, name, position, height, play_grade FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []player.Player
	for rows.Next() {
		var id, name, position string
		var height float32
		var play_grade uint8

		if err := rows.Scan(&id, &name, &position, &height, &play_grade); err != nil {
			return nil, err
		}

		players = append(players, player.Player{
			ID: id,
			Name: name,
			Position: position,
			Height: height,
			PlayGrade: play_grade,
		})
	}

	return players, nil
}
