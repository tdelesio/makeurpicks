package dao

import "makeurpicks/league/model"

type LeagueRepository interface {
	CreateLeague(player model.League) (model.League, error)
	UpdateLeague(player model.League) (model.League, error)
	GetLeague(id string) (model.League, error)
	GetLeagueByLeagueName(leaguename string) (model.League, error)
	GetLeagues() ([]model.League,error)

}