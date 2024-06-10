package usecase

type InputCreatePlayer struct {
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Height    float32 `json:"height"`
	PlayGrade uint8   `json:"play_grade"`
}

type OutputPlayer struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Position  string  `json:"position"`
	Height    float32 `json:"height"`
	PlayGrade uint8   `json:"play_grade"`
}
