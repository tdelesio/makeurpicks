package dao

import "makeurpicks/player/model"

type PlayerRepository interface {
	CreatePlayer(player model.Player) (model.Player, error)
	UpdatePlayer(player model.Player) (model.Player, error)
	GetPlayer(username string) (model.Player, error)
	GetPlayers() ([]model.Player,error)

}