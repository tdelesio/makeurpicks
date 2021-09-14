package dao

import "makeurpicks/game/model"

type GameRepository interface {

	CreateGame(game model.Game) (model.Game,error)
	UpdateGame(game model.Game) (model.Game,error)
	GetGame(gameid string) (model.Game,error)
	GetGamesByWeek(weekid string) ([]model.Game, error)
	GetGamesByWeekOrderByGameStart(weekid string) ([]model.Game, error)
	//FindByWeekIdAndFavIdAndDogId(weekid string, favid string, dogid string) (model.Game, error)
}

