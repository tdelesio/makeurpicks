package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TeamName string
	City string
	ShortName string
	Theme string
	FeedName string
	LeagueType string
}
