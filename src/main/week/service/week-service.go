package service

import (
	"makeurpicks/week/dao"
	"makeurpicks/week/model"
)

type WeekService struct {
	WeekRepository dao.WeekRepository
}

func (w WeekService) GetWeeksBySeason(seasonid string)([]model.Week, error) {
	return w.WeekRepository.GetWeeksBySeason(seasonid)
}

func (w WeekService) CreateWeek(week model.Week)(model.Week,error) {
	return w.WeekRepository.CreateWeek(week)
}
