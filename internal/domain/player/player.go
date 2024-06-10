package player

import (
	"errors"

	"github.com/google/uuid"
)

type Player struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Height    float32 `json:"height"`
	PlayGrade uint8   `json:"play_grade"`
}

func NewPlayer(name, position string, height float32, playGrade uint8) (*Player, error) {
	if err := isHeightValid(height); err != nil {
		return nil, err
	}

	if err := isPlayGradeValid(playGrade); err != nil {
		return nil, err
	}

	return &Player{
		ID:        uuid.NewString(),
		Name:      name,
		Position:  position,
		Height:    height,
		PlayGrade: playGrade,
	}, nil
}

func isHeightValid(height float32) error {
	if height > 2.2 || height < 1.5 {
		return errors.New("invalid height")
	}

	return nil
}

func isPlayGradeValid(playGrade uint8) error {
	if playGrade > 5 {
		return errors.New("invalid play grade")
	}

	return nil
}
