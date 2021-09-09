package model

type Team struct {
	ID string `json:"_id"`
	TeamName string
	City string
	ShortName string
	Theme string
	FeedName string
	LeagueType string
}
