package game

type GameService struct {
	GameRepository GameRepository
}

//func (gameService *GameService) CreateGame(game model.Game) (err error) {
//
//	//gameService.ValidateGame(game)
//}

//func (gameService *GameService) ValidateGame (game model.Game) (err error) {
//	//if game==nil {
//	//	err = errors.New("game service -> invalid input, game is null")
//	//	return
//	//}
//
//	if "" == game.Favid {
//		err = errors.New("game service -> invalid input, favid is null")
//	}
//
//	if "" == game.Dogid {
//		err = errors.New("game service -> invalid input, dogid is null")
//	}
//
//	if "" == game.Weekid {
//		err = errors.New("game service -> invalid input, weekid is null")
//	}
//
//	//if check.NotNil(game.Gamestart){
//	//	err = errors.New("game service -> invalid input, gameStart is null")
//	//}
//
//	if game.Dogid == game.Favid {
//		err = errors.New("game service -> invalid input, team cannot play itself")
//	}
//}
//
//func (gameService *GameService) CreateGame (game model.Game) (err error) {
//	gameService.ValidateGame(game)
//
//	model.Team fav := team.TeamRepository.GetTeam(game.Favid)
//	model.Team dog := team.TeamRepository.GetTeam(game.Dogid)
//
//	model.Game.Id = uuid.NewString()
//	model.Game.Dog
//
//	{
//		Id:           uuid.NewString(),
//		Spread:       0,
//		Seasonid:     "",
//		Favid:        "",
//		Dogid:        "",
//		Weekid:       "",
//		Favscore:     0,
//		Dogscore:     0,
//		Favhome:      false,
//		Gamestart:    time.Time{},
//		Favfullname:  "",
//		Dogfullname:  "",
//		Dogshortname: "",
//		Favshortname: "",
//	}
	//Team fav = teamService.getTeam(game.getFavId());
	//Team dog = teamService.getTeam(game.getDogId());
	//game.setDogShortName(dog.getShortName());
	//game.setFavShortName(fav.getShortName());
	//game.setFavFullName(fav.getFullTeamName());
	//game.setDogFullName(dog.getFullTeamName());

//}

//
//public Game createGame(Game game)
//{
//validateGame(game);
//
//game.generateId();
//
//Team fav = teamService.getTeam(game.getFavId());
//Team dog = teamService.getTeam(game.getDogId());
//game.setDogShortName(dog.getShortName());
//game.setFavShortName(fav.getShortName());
//game.setFavFullName(fav.getFullTeamName());
//game.setDogFullName(dog.getFullTeamName());
//
//return gameRepository.save(game);
//}
//
//public Game updateGame(Game game)
//{
//validateGame(game);
//
////allow only certain fields to be updated
//Game gameFromDS = gameRepository.findOne(game.getId());
//gameFromDS.setGameStart(game.getGameStart());
//gameFromDS.setSpread(game.getSpread());
//gameFromDS.setFavHome(game.isFavHome());
//gameFromDS.setFavScore(game.getFavScore());
//gameFromDS.setDogScore(game.getDogScore());
//
//return gameRepository.save(gameFromDS);
//}
//
//public Game updateGameScore(Game game)
//{
//Game gameFromDS = gameRepository.findOne(game.getId());
//gameFromDS.setFavScore(game.getFavScore());
//gameFromDS.setDogScore(game.getDogScore());
//return gameRepository.save(gameFromDS);
//}
//
//public List<Game> getGamesByWeek(String weekId)
//{
//List<Game> games = gameRepository.findByWeekId(weekId);
//
//Collections.sort(games, (g1, g2)-> g1.getGameStart().compareTo(g2.getGameStart()));
//return games;
////		return gameRepository.findByWeekIdOrderByGameStart(weekId);
//}
//
//public Game getGameById(String gameId)
//{
//Game game = gameRepository.findOne(gameId);
//return game;
//}
//
//
//
//
//public NFLWeek loadFromNFL()
//{
//RestTemplate restTemplate = new RestTemplate();
//restTemplate.getMessageConverters().add(new StringHttpMessageConverter());
//return restTemplate.getForObject("http://www.nfl.com/liveupdate/scorestrip/ss.json", NFLWeek.class);
//}
//
//public void updateScoreFromNFL(String weekId)
//{
//NFLWeek nflWeek = loadFromNFL();
//
//
//for (NFLGame nflGame:nflWeek.getGms())
//{
//Team home = teamService.getTeam(nflGame.getH());
//Team away = teamService.getTeam(nflGame.getV());
//
//Game game = gameRepository.findByWeekIdAndFavIdAndDogId(weekId, home.getId(), away.getId());
//if (game == null)
//{
//game = gameRepository.findByWeekIdAndFavIdAndDogId(weekId, away.getId(), home.getId());
//if (game == null)
//{
//log.debug("No game found for "+home.getFullTeamName()+ " - "+away.getFullTeamName());
////no game found, continue;
//continue;
//}
//else
//{
//game.setFavScore(nflGame.getVs());
//game.setDogScore(nflGame.getHs());
//}
//}
//else
//{
//game.setFavScore(nflGame.getHs());
//game.setDogScore(nflGame.getVs());
//}
//
//updateGameScore(game);
//}
//}
//
//public void autoSetupWeek(String seasonId)
//{
//if (seasonId == null || "".equals(seasonId))
//throw new GameValidationException(GameExceptions.SEASON_ID_IS_NULL);
//List<Week> weeks = weekService.getWeeksBySeason(seasonId);
//
//NFLWeek nflWeek = loadFromNFL();
//int weekNumber = nflWeek.getW();
//String foundWeekId = null;
//
//for (Week week: weeks)
//{
//if (week.getWeekNumber() == weekNumber)
//{
//foundWeekId = week.getId();
//break;
//}
//}
//
//if (foundWeekId == null)
//{
////week not found so we need to create one
//Week newWeek = new WeekBuilder().withWeekNumber(weekNumber).withSeasonId(seasonId).build();
//Week week = weekService.createWeek(newWeek);
//
////now that we have a week setup, let's create the games
//for (NFLGame nflGame:nflWeek.getGms())
//{
//Team home = teamService.getTeam(nflGame.getH());
//Team away = teamService.getTeam(nflGame.getV());
//String day = nflGame.getD();
//String time = nflGame.getT();
//ZonedDateTime gameStart = ZonedDateTime.now();
//StringTokenizer stringTokenizer = new StringTokenizer(time, ":");
//String hour = stringTokenizer.nextToken();
//String min = stringTokenizer.nextToken();
//gameStart = gameStart
//.withHour(Integer.parseInt(hour)+12)
//.withMinute(Integer.parseInt(min))
//.withSecond(0)
//.withNano(0);
//
////based on running on wed
//int dayOffSet = 0;
//if (day.equals("Thu"))
//dayOffSet++;
//else if (day.equals("Sat"))
//dayOffSet+=3;
//else if (day.equals("Sun"))
//dayOffSet+=4;
//else if (day.equals("Mon"))
//dayOffSet+=5;
//
////monday is 1, tues is 2, wed is 3, thur is 4
//int currentDay = gameStart.getDayOfWeek().getValue();
//if (currentDay == 2)
//dayOffSet++;
//else if (currentDay == 4)
//dayOffSet--;
//else if (currentDay == 5)
//dayOffSet-=2;
//
//gameStart = gameStart.plusDays(dayOffSet);
//Game game = new GameBuilder()
//.withFavId(home.getId())
//.withDogId(away.getId())
//.withWeekId(week.getId())
//.withGameStartTime(gameStart)
//.build();
//
//createGame(game);
//}
//}

//}