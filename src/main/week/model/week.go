package model

type Week struct {
	ID string `json:"id" bson:"_id,omitempty"`

	//Season model.Season
	//Games []model2.Game
	SeasonId string `json:"seasonid" bson:"seasonid,omitempty"`
	WeekNumber int
}