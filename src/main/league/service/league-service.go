package service

import (
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
	"makeurpicks/league/dao"
	"makeurpicks/league/model"
	model2 "makeurpicks/player/model"
	"makeurpicks/player/service"
)

func GetLeagueTypes()([]string) {
	return []string{PICKEM, SUICIDE}
}

const PICKEM = `pickem`
const SUICIDE = `suicide`

type LeagueService struct {
	LeagueRepository dao.LeagueRepository
	PlayerService *service.PlayerService
}

func validateLeague(league model.League) {
	v := validator.New()
	v.Struct(league)
}

func (l LeagueService) CreateLeague (league model.League)(model.League,error){
	validateLeague(league)

	return l.LeagueRepository.CreateLeague(league)
}

func (l LeagueService) UpdateLeague(league model.League)(model.League,error) {
	validateLeague(league)

	return l.LeagueRepository.UpdateLeague(league)
}

func (l LeagueService) GetLeague(id string) (model.League, error) {
	return l.LeagueRepository.GetLeague(id)
}

func (l LeagueService) GetLeagueByLeagueName(leaguename string) (model.League, error) {
	return l.GetLeagueByLeagueName(leaguename)
}

func (l LeagueService) GetPlayersInLeague(id string) ([]model2.Player, error) {

	league, err := l.GetLeague(id)

	if err != nil{
		panic(err)
	}

	var players []model2.Player
	for _, playerid := range league.Players {
		player,err := l.PlayerService.GetPlayer(playerid)
		if err != nil {
			panic(err)
		}

		players = append(players, player)
	}

	return players, err
}

func (l LeagueService) JoinLeague(leagueid string, playerid string, password string) error{
	league,err := l.LeagueRepository.GetLeague(leagueid)
	if err != nil {
		panic(err)
	}

	if league.Password != password {
		return errors.New("Password does not match League Password.  Please reach out to the league admin")
	}

	league.Players = append(league.Players, playerid)
	l.LeagueRepository.UpdateLeague(league)

	player,err := l.PlayerService.GetPlayer(playerid)
	if err != nil {
		panic(err)
	}

	player.Leagues = append(player.Leagues, leagueid)
	l.PlayerService.UpdatePlayer(player)

	return nil

}

func (l LeagueService) LeaveLeague(leagueid string, playerid string) error{
	league,err := l.LeagueRepository.GetLeague(leagueid)
	if err != nil {
		panic(err)
	}

	player,err := l.PlayerService.GetPlayer(playerid)
	if err != nil {
		panic(err)
	}

	leagueids := league.Players
	var updatedleagueids []string

	for _, lid := range leagueids {
		if lid != leagueid {
			updatedleagueids = append(updatedleagueids, lid)
		}
	}
	player.Leagues = updatedleagueids
	l.PlayerService.UpdatePlayer(player)

	playerids := player.Leagues
	var updatedplayerids [] string

	for _, pid := range playerids {
		if pid != playerid {
			updatedplayerids = append(updatedplayerids, pid)
		}
	}
	league.Players = updatedplayerids
	l.LeagueRepository.UpdateLeague(league)

	return nil
}