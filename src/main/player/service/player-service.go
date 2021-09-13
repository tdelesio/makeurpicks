package service

import (
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	model2 "makeurpicks/league/model"
	"makeurpicks/league/service"
	"makeurpicks/player/dao"
	"makeurpicks/player/model"
)

const HASH = "makeyourpicksisthebest"

type PlayerService struct {
	PlayerRepository dao.PlayerRepository
	LeagueService *service.LeagueService
}

func ValidatePlayer(player model.Player) {
	player.Init()
	v := validator.New()
	v.Struct(player)

	return
}

func (dao PlayerService)GetPlayer(id string)(model.Player, error) {
	return dao.PlayerRepository.GetPlayer(id)
}

func (p PlayerService)UpdatePlayer(player model.Player)(model.Player, error) {
	return p.PlayerRepository.UpdatePlayer(player)
}

func (p PlayerService)GetPlayerByUsername(username string) (model.Player, error) {
	return p.PlayerRepository.GetPlayerByUsername(username)
}

func (dao PlayerService)Login(username string, password string) (bool, error) {
	player, err := dao.GetPlayer(username)

	return CheckPasswordHash(player.Password, HASH), err
}

func (p PlayerService)UpdatePassword(username string, password string)(error) {
	hashedPassword, err := HashPassword(password)

	if err!= nil {
		panic(err)
	}

	player, err := p.GetPlayer(username)
	player.Password=hashedPassword

	player, err = p.PlayerRepository.UpdatePlayer(player)
	return err

}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (p PlayerService)RegisterPlayer(player model.Player)(model.Player, error) {
	ValidatePlayer(player)

	_,err := p.GetPlayer(player.Username)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return p.PlayerRepository.CreatePlayer(player)
		}
	} else {
		panic(err)
	}

	return player, err

}

func (p PlayerService) GetLeaguesForPlayer (id string) ([]model2.League, error) {
	var leagues []model2.League

	player, err := p.GetPlayer(id)
	if err != nil {
		panic(err)
	}

	for _, leagueid := range player.Leagues {
		league, err := p.LeagueService.GetLeague(leagueid)
		if err != nil {
			panic(err)
		}

		leagues = append(leagues, league)
	}

	return leagues, err
}
