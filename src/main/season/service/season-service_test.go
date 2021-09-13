package service

import (
	"github.com/stretchr/testify/assert"
	"makeurpicks/league/service"
	"makeurpicks/season/dao"
	"testing"
)

func TestSeasonService_GetCurrentSeasons(t *testing.T) {

	seasonService := SeasonService{SeasonRepository: dao.SeasonRepositoryTestDummy{}}

	seasons, err := seasonService.GetCurrentSeasons()

	assert.Equal(t, 2, len(seasons))
	assert.Nil(t, err)

	foundPickem := false
	foundSuicide := false

	for _, s := range seasons {
		if s.LeagueType == service.PICKEM {
			foundPickem = true
		}
		if s.LeagueType == service.SUICIDE {
			foundSuicide = true
		}
	}

	assert.True(t, foundSuicide)
	assert.True(t, foundPickem)
}
