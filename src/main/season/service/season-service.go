package service

import (
	"makeurpicks/league/service"
	"makeurpicks/season/dao"
	"makeurpicks/season/model"
	"time"
)

type SeasonService struct {
	SeasonRepository dao.SeasonRepository
}

func (s SeasonService)CreateSeason(season model.Season)(model.Season, error) {
	return s.SeasonRepository.CreateSeason(season)
}

func (s SeasonService)DeleteSeason(id string) {
	s.SeasonRepository.DeleteSeason(id)
}

func (s SeasonService)UpdateSeason(season model.Season)(model.Season, error) {
	return s.SeasonRepository.UpdateSeason(season)
}

func (s SeasonService)GetCurrentSeasons()([]model.Season, error) {

	var err error
	var currentSeasons []model.Season
	leagueTypes := service.GetLeagueTypes()
	t := time.Now()
	currentYear := t.Year()
	for _, leagueType := range leagueTypes {

		seasons, serr := s.SeasonRepository.GetSeasonsByLeagueType(leagueType)
		if err != nil {
			err = serr
		}
		for _, season := range seasons {
			if season.StartYear == currentYear {
				currentSeasons = append(currentSeasons,season)
			}

		}

	}

	return currentSeasons, err
}