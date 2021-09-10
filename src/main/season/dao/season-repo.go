package dao

import (
	"makeurpicks/season/model"
)

type SeasonRepository interface {

	GetSeasonsByLeagueType(leagyetype string) ([]model.Season, error)
	CreateSeason(season model.Season) (model.Season,error)
	UpdateSeason(season model.Season) (model.Season,error)
	DeleteSeason(id string) error
}






