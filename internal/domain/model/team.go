package model

type Team struct {
	Id    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name  string `gorm:"size:255 unique"`
	Users []User `gorm:"many2many:user_teams;constraint:OnDelete:CASCADE;"`
}

type UserTeam struct {
	UserId string `gorm:"primaryKey"`
	TeamId string `gorm:"primaryKey"`
	State  string `gorm:"size:255"` // "INVITED" or "MEMBER"
}
