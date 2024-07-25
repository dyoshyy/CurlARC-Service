package model

import (
	"time"

	"gorm.io/datatypes"
)

type Record struct {
	Id       string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TeamId   string
	Place    string `gorm:"size:255"`
	Date     time.Time
	Team     Team           `gorm:"foreignKey:TeamId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EndsData datatypes.JSON `gorm:"type:json"`
}

type DataPerEnd struct {
	Index int
	Score int
	Shots []Shot
}

type Shot struct {
	Index       int
	Type        string
	SuccessRate float64
	Shooter     string
	Stones      Stones
}

type Stones struct {
	FriendStones []Coordinate
	EnemyStones  []Coordinate
}

type Coordinate struct {
	Index int
	R     float64
	Theta float64
}
