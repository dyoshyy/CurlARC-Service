package infra

import (
	"time"

	"gorm.io/datatypes"
)

type Team struct {
	Id      string   `gorm:"primaryKey"`
	Name    string   `gorm:"type:varchar(100)"`
	Records []Record `gorm:"foreignKey:TeamId"`
	Users   []User   `gorm:"many2many:user_teams;"`
}

// define the struct for the database
type User struct {
	Id    string `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"uniqueIndex;type:varchar(100)"`
	Teams []Team `gorm:"many2many:user_teams;"`
}

type Record struct {
	Id            string         `gorm:"type:uuid;primaryKey"`
	TeamId        string         `gorm:"foreignKey:TeamId"`
	Result        string         `gorm:"type:varchar(10)"`
	EnemyTeamName string         `gorm:"type:varchar(255)"`
	Place         string         `gorm:"type:varchar(255)"`
	Date          time.Time      `gorm:"type:timestamp"`
	EndsDataJSON  datatypes.JSON `gorm:"type:json"`
	IsPublic      bool           `gorm:"type:boolean"`
}