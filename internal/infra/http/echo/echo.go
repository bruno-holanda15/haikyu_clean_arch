package echo

import (
	// "encoding/json"

	"fmt"
	"haikyu_game/internal/domain/player"
	"haikyu_game/internal/domain/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handlers(repository player.RepositoryInterface) *echo.Echo {
	e := echo.New()

	playerHandler := NewPlayerHandler(repository)
	e.POST("/createPlayer", playerHandler.Create)

	return e
}

type PlayerHandler struct {
	repository player.RepositoryInterface
}

func NewPlayerHandler(repository player.RepositoryInterface) *PlayerHandler {
	return &PlayerHandler{
		repository: repository,
	}
}

func (p *PlayerHandler) Create(c echo.Context) error {
	var input usecase.InputCreatePlayer
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	createUseCase := usecase.NewCreatePlayerUseCase(p.repository)
	output, err := createUseCase.Execute(input)
	if err != nil {
		errMsg := fmt.Sprintf("error creating player - %v", err.Error())
		c.String(http.StatusInternalServerError, errMsg)
		return err
	}

	return c.JSON(http.StatusCreated, output)
}
