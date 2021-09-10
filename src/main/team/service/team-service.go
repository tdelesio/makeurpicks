package service

import (
	"fmt"
	"makeurpicks/team/dao"
	"makeurpicks/team/model"
)


type TeamService struct {
	TeamRepository dao.TeamRepository
}

func (s *TeamService)createTeam(team model.Team)(*model.Team, error) {
	return s.TeamRepository.CreateTeam(team)
}

func (s *TeamService)GetAllTeams(leagueType string)(*[]model.Team, error) {
	return s.TeamRepository.GetTeams(leagueType)
}

func (s *TeamService)GetTeam(id string) (model.Team, error) {
	return s.TeamRepository.GetTeam(id)
}

func (s *TeamService)BuildTeamMap(leaguetype string) (map[string]model.Team) {
	teamMap:= make(map[string]model.Team)

	teams,err := s.GetAllTeams(leaguetype)
	if err != nil {
		// Checks if the movie was not found
		panic(err)
	}

	for _, s := range *teams {
		teamMap[s.ID.String()] = s
	}

	return teamMap
}
func (s *TeamService)CreateAllTeams(leagueType string)  {

	teams,err := s.GetAllTeams(leagueType)
	if err != nil {
		// Checks if the movie was not found
		panic(err)
	}
	numTeams := len(*teams)
	if numTeams == 0 {
		fmt.Println("No Teams Found, let's create them")
		s.CreateAllTeams(leagueType)


		team := model.Team{
			TeamName:   "Cardinals",
			City:       "Arizona",
			ShortName:  "ARI",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Bills",
			City:       "Buffalo",
			ShortName:  "BUF",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Dolphins",
			City:       "Miami",
			ShortName:  "MIA",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Patriots",
			City:       "New England",
			ShortName:  "NE",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Jets",
			City:       "New York",
			ShortName:  "NYJ",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Ravens",
			City:       "Baltimore",
			ShortName:  "BAL",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Bengals",
			City:       "Cincinnati",
			ShortName:  "CIN",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Browns",
			City:       "Cleveland",
			ShortName:  "CLE",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Panthers",
			City:       "Carolina",
			ShortName:  "CAR",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Steelers",
			City:       "Pittsburg",
			ShortName:  "PIT",
			LeagueType: leagueType,
		}
		s.createTeam(team)
		team = model.Team{
			TeamName:   "Texans",
			City:       "Houston",
			ShortName:  "HOU",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Colts",
			City:       "Indianapolis",
			ShortName:  "IND",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Jaguars",
			City:       "Jacksinville",
			ShortName:  "JAC",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Titans",
			City:       "Tennessee",
			ShortName:  "TEN",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Broncis",
			City:       "Denver",
			ShortName:  "DEN",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Chiefs",
			City:       "Kansis City",
			ShortName:  "KC",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Raiders",
			City:       "Las Vegas",
			ShortName:  "LVR",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Chargers",
			City:       "Los Angels",
			ShortName:  "LAC",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Cowboys",
			City:       "Dallas",
			ShortName:  "DAL",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Giants",
			City:       "New York",
			ShortName:  "NYG",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Eagles",
			City:       "Philladelpha",
			ShortName:  "PHL",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Football Team",
			City:       "Washington",
			ShortName:  "WAS",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Bears",
			City:       "Chicago",
			ShortName:  "CHI",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Lions",
			City:       "Detroit",
			ShortName:  "DET",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Packers",
			City:       "Green Bay",
			ShortName:  "GB",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Vikings",
			City:       "Minnesota",
			ShortName:  "MIN",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Rams",
			City:       "St. Loius",
			ShortName:  "STL",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Falcons",
			City:       "Atlanta",
			ShortName:  "ATL",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Bucineers",
			City:       "Tampa Bay",
			ShortName:  "TB",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Saints",
			City:       "New Orleans",
			ShortName:  "NO",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "49ers",
			City:       "San Fransico",
			ShortName:  "SF",
			LeagueType: leagueType,
		}
		s.createTeam(team)

		team = model.Team{
			TeamName:   "Seahawks",
			City:       "Seattle",
			ShortName:  "SEA",
			LeagueType: leagueType,
		}
		s.createTeam(team)
	} else if numTeams != 32 {
		panic("Number of teams in the database is not equal to 32, requires investigation")
	} else {
		fmt.Println("Teams are already in the DB, continue")
	}

}

