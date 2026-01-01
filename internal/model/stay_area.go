package model

import "github.com/google/uuid"

type StayArea struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name     string    `gorm:"not null"`
	Capacity int       `gorm:"not null"`
}
