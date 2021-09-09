package team

import (
	"github.com/stretchr/testify/assert"
	"makeurpicks/model"
	"testing"
)


func TestCreateTeam(t *testing.T) {

	teamService := TeamService{
		TeamRepository: &TeamRepositoryTestDummy{},
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




}

func TestGetTeam(t *testing.T) {

	teamService := TeamService{
		TeamRepository: &TeamRepositoryTestDummy{},
	}

	teams, err := teamService.GetTeam("1")
	assert.NotNil(t, teams)
	assert.Nil(t, err)

	teams, err = teamService.GetTeam("2")
	assert.Nil(t, teams)
	assert.NotNil(t, err)
}

func TestBuildTeamMap(t *testing.T) {

}

func TestCreateAllTeams(t *testing.T) {

}