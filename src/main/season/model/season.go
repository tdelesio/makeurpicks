package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Season struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LeagueType string
	StartYear int
	EndYear int
}
