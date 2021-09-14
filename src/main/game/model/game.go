package model

import (
	model2 "makeurpicks/team/model"
	"time"
)

type Game struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Spread   float32 `json:"spread"`
	Fav model2.Team
	Dog model2.Team
	//Favid  string `json:"favidid"`
	//Dogid  string `json:"dogidd"`

	Seasonid string `json:"seasonid"`
	Weekid string `json:"weekidid"`

	Favscore int  `json:"favscore"`
	Dogscore  int       `json:"dogscore"`
	Favhome   bool      `json:"favhome"`
	Gamestart time.Time `json:"gamestart"`

	//Favfullname  string `json:"favfullname"`
	//Dogfullname  string `json:"dogfullname"`
	//Dogshortname string `json:"dogshortname"`
	//Favshortname string `json:"favshortname"`
}