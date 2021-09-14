package service

import (
	"gopkg.in/go-playground/validator.v9"
	"makeurpicks/game/dao"
	"makeurpicks/game/model"
	team "makeurpicks/team/service"
)

type GameService struct {
	GameRepository dao.GameRepository
	TeamService team.TeamService
}

func validateGame(game model.Game) {
	v := validator.New()
	v.Struct(game)
}

func (g GameService) CreateGame (game model.Game, favid string, dogid string) (model.Game, error) {
	validateGame(game)

	fav,err := g.TeamService.GetTeam(favid)
	if err != nil {
		panic(err)
	}

	dog,err := g.TeamService.GetTeam(dogid)
	if err != nil {
		panic(err)
	}

	game.Fav = fav
	game.Dog = dog

	return g.GameRepository.CreateGame(game)
}

func (g GameService) UpdateGame (game model.Game) (model.Game,error) {
	validateGame(game)

	gameFromDS,err := g.GameRepository.GetGame(game.ID)
	if err != nil {
		panic(err)
	}
	gameFromDS.Gamestart = game.Gamestart
	gameFromDS.Spread = game.Spread
	gameFromDS.Favhome = game.Favhome
	gameFromDS.Favscore = game.Favscore
	gameFromDS.Dogscore = game.Dogscore

	return g.GameRepository.UpdateGame(gameFromDS)

}

func (g GameService) GetGame(gameid string) (model.Game,error) {
	return g.GameRepository.GetGame(gameid)
}

func (g GameService) GetGamesByWeek(weekid string) ([]model.Game, error) {
	return g.GameRepository.GetGamesByWeekOrderByGameStart(weekid)
}