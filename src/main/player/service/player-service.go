package service

import (
	"makeurpicks/player/dao"
)

type PlayerService struct {
	PlayerRepository dao.PlayerDaoMongo
}

//func (p *PlayerService)GetPlayer(username string)(*model.Player, error) {
//	return p.PlayerRepository.
//}