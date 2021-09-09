package team

import (
	"go.mongodb.org/mongo-driver/mongo"
	"makeurpicks/model"
)

type TeamRepository interface {

	//GetTeam(gameid string) (model.Team, error)
	GetTeams(leaguetype string) ([]*model.Team,error)
	CreateTeam(team model.Team) (*mongo.InsertOneResult,error)
}