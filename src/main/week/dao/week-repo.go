package dao

import "makeurpicks/week/model"

type WeekRepository interface {
	GetWeeksBySeason(seasonid string) ([]model.Week, error)
	CreateWeek(week model.Week) (model.Week,error)
}

