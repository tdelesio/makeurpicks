package model

import "go.mongodb.org/mongo-driver/bson/primitive"

const ADMIN = 1
const LEAGUE_ADMIN = 2
const USER = 3

const ACTIVE = 1
const UNPAID = 2
const DISABLED = 3

type Player struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name" validate:"required,name"`
	Address string `json:"address" validate:"required,address"`
	PlayerStatus int
	MemberLevel int

	Leagues []string

}

func (p *Player) Init() {
	if p.PlayerStatus == 0 {
		p.PlayerStatus = ACTIVE
	}
	if p.MemberLevel == 0 {
		p.MemberLevel = USER
	}
}
