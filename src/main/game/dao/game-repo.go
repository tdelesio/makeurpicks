package dao

import "makeurpicks/game/model"

type GameRepository interface {

	FindByWeekId(weekid string) ([]model.Game, error)
	FindByWeekIdOrderByGameStart(weekid string) ([]model.Game, error)
	FindByWeekIdAndFavIdAndDogId(weekid string, favid string, dogid string) (model.Game, error)
}