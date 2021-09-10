package model

const ADMIN = 1
const LEAGUE_ADMIN = 2
const USER = 3

const ACTIVE = 1
const UNPAID = 2
const DISABLED = 3

type Player struct {
	Username string `json:"username" bson:"_id,omitempty"`
	Password string
	Email string
	Name string
	Address string
	PlayerStatus int
	MemberLevel int

}
