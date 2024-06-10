package player

type RepositoryInterface interface {
	Save(player *Player) error
	ListAll() ([]Player, error)
}
