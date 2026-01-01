package model

import (
	"time"

	"github.com/google/uuid"
)

type SevaType struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description *string
	IsActive    bool `gorm:"default:true"`
	CreatedAt   time.Time
}

type Schedule struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProfileID  uuid.UUID `gorm:"type:uuid;not null;index"`
	Profile    Profile   `gorm:"foreignKey:ProfileID;references:ID"`
	VisitID    uuid.UUID `gorm:"type:uuid;not null;index"`
	Visit      Visit     `gorm:"foreignKey:VisitID;references:ID"`
	Date       time.Time `gorm:"not null"`
	SevaTypeID uuid.UUID `gorm:"type:uuid;not null;index"`
	SevaType   SevaType  `gorm:"foreignKey:SevaTypeID"`
	Location   *string   `gorm:"type:varchar(200)"`
	Notes      *string   `gorm:"type:text"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
