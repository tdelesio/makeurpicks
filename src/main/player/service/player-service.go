package service

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"makeurpicks/player/dao"
	"makeurpicks/player/model"
)

const HASH = "makeyourpicksisthebest"

type PlayerService struct {
	PlayerRepository dao.PlayerRepository
}

func ValidatePlayer(player model.Player) {
	player.Init()
	v := validator.New()
	v.Struct(player)

	return
}

func (dao PlayerService)GetPlayer(username string)(model.Player, error) {
	return dao.PlayerRepository.GetPlayer(username)
}

func (dao PlayerService)Login(username string, password string) (bool, error) {
	player, err := dao.GetPlayer(username)

	return CheckPasswordHash(player.Password, HASH), err
}

func (p PlayerService)updatePassword(username string, password string)(error) {
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

func (p PlayerService)registerPlayer(player model.Player)(model.Player, error) {
	ValidatePlayer(player)

	return p.PlayerRepository.CreatePlayer(player)
}