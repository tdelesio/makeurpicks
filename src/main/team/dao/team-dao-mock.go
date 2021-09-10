package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"makeurpicks/team/model"
)

type TeamRepositoryTestDummy struct {}

func (dao *TeamRepositoryTestDummy) CreateTeam(team model.Team) (*model.Team,error) {
	return &model.Team{
		ID:         primitive.NewObjectID(),
		TeamName:   team.TeamName,
		City:       team.City,
		ShortName:  team.ShortName,
		Theme:      team.Theme,
		FeedName:   team.FeedName,
		LeagueType: team.LeagueType,
	}, nil
}

func (dao *TeamRepositoryTestDummy) GetTeam(id string) (model.Team, error) {


	switch id {
	case "1":
		return model.Team{
			ID:         primitive.NewObjectID(),
			TeamName:   "Team",
			City:       "City",
			ShortName:  "TST",
			Theme:      "",
			FeedName:   "",
			LeagueType: "pickem",
		}, nil
	default:
		err := mongo.ErrNoDocuments
		return model.Team{}, err

	}
}

func (dao *TeamRepositoryTestDummy) GetTeams(leagueType string) (*[]model.Team,error) {

	teams := make([]model.Team, 2)
	teams = append(teams, model.Team{
		ID:         primitive.NewObjectID(),
		TeamName:   "Team 1",
		City:       "City",
		ShortName:  "TSS!",
		LeagueType: "pickem",
	}, model.Team{
		ID:         primitive.NewObjectID(),
		TeamName:   "Team 2",
		City:       "City 2",
		ShortName:  "TST",
		Theme:      "",
		FeedName:   "",
		LeagueType: "pickem",
	})

	return &teams, nil
}