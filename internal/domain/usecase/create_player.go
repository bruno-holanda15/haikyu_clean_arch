package usecase

import "haikyu_game/internal/domain/player"

type CreatePlayerUsecase struct {
	repo player.RepositoryInterface
}


func NewCreatePlayerUseCase(repo player.RepositoryInterface) *CreatePlayerUsecase {
	return &CreatePlayerUsecase{
		repo: repo,
	}
}

func (c *CreatePlayerUsecase) Execute(input InputCreatePlayer) (OutputPlayer, error) {
	p, err := player.NewPlayer(input.Name, input.Position, input.Height, input.PlayGrade)
	if err != nil {
		return OutputPlayer{}, err
	}

	err = c.repo.Save(p)
	if err != nil {
		return OutputPlayer{}, err 
	}

	return OutputPlayer{p.ID, p.Name, p.Position, p.Height, p.PlayGrade}, nil
}