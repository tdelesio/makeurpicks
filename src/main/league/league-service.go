package league

func GetLeagueTypes()([]string) {
	return []string{PICKEM, SUICIDE	}
}

const PICKEM = `pickem`
const SUICIDE = `suicide`