package service

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"makeurpicks/team/dao"
	"makeurpicks/team/model"
	"testing"
)


func TestCreateTeam(t *testing.T) {

	teamService := TeamService{
		TeamRepository: &dao.TeamRepositoryTestDummy{},
	}

	input := model.Team{
		TeamName:   "Team Name",
		City:       "City",
		ShortName:  "SRT",
		Theme:      "",
		FeedName:   "",
		LeagueType: "pickem",
	}
	team, err := teamService.createTeam(input)

	assert.Equal(t, team.TeamName, input.TeamName)
	assert.Equal(t, team.City, input.City)
	assert.Equal(t, team.ShortName, input.ShortName)
	assert.Equal(t, team.LeagueType, input.LeagueType)
	assert.Equal(t, team.Theme, input.Theme)
	assert.Equal(t, team.FeedName, input.FeedName)
	assert.Nil(t, err)


}

func TestGetAllTeams(t *testing.T) {

	teamService := TeamService{
		TeamRepository: &dao.TeamRepositoryTestDummy{},
	}

	teams, err := teamService.GetAllTeams("pickem")
	assert.Nil(t, err)
	assert.Equal(t, len(*teams), 2)


}

func TestGetTeam(t *testing.T) {

	teamService := TeamService{
		TeamRepository: &dao.TeamRepositoryTestDummy{},
	}

	teams, err := teamService.GetTeam("1")
	assert.NotNil(t, teams)
	assert.Nil(t, err)

	teams, err = teamService.GetTeam("2")
	assert.Equal(t, err, mongo.ErrNoDocuments)
}

func TestBuildTeamMap(t *testing.T) {


	teamService := TeamService{
		TeamRepository: &dao.TeamRepositoryTestDummy{},
	}

	team := teamService.BuildTeamMap("pickem")
	assert.Equal(t, len(team), 32)

}