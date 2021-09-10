package model

import "time"

type Game struct {
	ID string `json:"id" bson:"_id,omitempty"`
	Spread   float32 `json:"spread"`
	Seasonid string `json:"seasonid"`
	Favid  string `json:"favidid"`
	Dogid  string `json:"dogidd"`
	Weekid string `json:"weekidid"`

	Favscore int  `json:"favscore"`
	Dogscore  int       `json:"dogscore"`
	Favhome   bool      `json:"favhome"`
	Gamestart time.Time `json:"gamestart"`

	Favfullname  string `json:"favfullname"`
	Dogfullname  string `json:"dogfullname"`
	Dogshortname string `json:"dogshortname"`
	Favshortname string `json:"favshortname"`
}