package model

type Week struct {
	ID string `json:"id" bson:"_id,omitempty"`
	seasonId string
	weekNumber int
}