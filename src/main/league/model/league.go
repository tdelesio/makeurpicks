package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type League struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LeagueName string `json:"leagueName" validate:"required,password"`
	PaidFor int
	Money bool
	Free bool
	Active bool
	Password string
	Spreads bool
	DoubleEnabled bool
	EntryFree float32
	WeeklyFee float32
	FirstPlacePercent int
	SecondPlacePercent int
	ThirdPlacePercent int
	FourthPlacePercent int
	FifthPlacePercent int
	DoubleType int
	Banker bool
	SeasonID primitive.ObjectID
	Admin string

	Players []string
}

func (l *League)init() {

}