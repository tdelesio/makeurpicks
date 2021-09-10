package dao

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"makeurpicks/league"
	"makeurpicks/season/model"
	"time"
)

type SeasonRepositoryTestDummy struct {}

func (dao *SeasonRepositoryTestDummy) GetSeasonsByLeagueType(leagyetype string) (*[]model.Season, error) {

	t := time.Now()
	currentYear := t.Year()
	followingYear := currentYear+1

	var seasons []model.Season

	switch leagyetype {
		case league.PICKEM:
			seasons = append(seasons, model.Season{
			ID:         primitive.NewObjectID(),
			LeagueType: league.PICKEM,
			StartYear:  currentYear,
			EndYear:    followingYear,
		}, model.Season{
			ID:         primitive.NewObjectID(),
			LeagueType: league.PICKEM,
			StartYear:  currentYear-1,
			EndYear:    followingYear-1,
		})
	case league.SUICIDE:
		seasons = append(seasons, model.Season{
			ID:         primitive.NewObjectID(),
			LeagueType: league.SUICIDE,
			StartYear:  currentYear,
			EndYear:    followingYear,
		})
	default:
		seasons = append(seasons, model.Season{})
	}




	return &seasons, nil
}

func (dao *SeasonRepositoryTestDummy) CreateSeason(season model.Season) (*model.Season, error) {
	return &model.Season{
		ID:         primitive.NewObjectID(),
		LeagueType: "pickem",
		StartYear:  2021,
		EndYear:    2022,
	}, nil
}

func (dao *SeasonRepositoryTestDummy) DeleteSeason(id string) error {
	return nil
}

func (dao SeasonRepositoryTestDummy) UpdateSeason(season model.Season) (*model.Season,error) {
	return &season, nil
}