package dao

import (
	"makeurpicks/team/model"
)

type TeamRepository interface {

	GetTeam(id string) (model.Team, error)
	GetTeams(leaguetype string) (*[]model.Team,error)
	CreateTeam(team model.Team) (*model.Team,error)
}