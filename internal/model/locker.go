package model

import (
	"time"

	"github.com/google/uuid"
)

type Locker struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	LockerNumber string    `gorm:"not null"`
	Section      string    `gorm:"not null"`
	IsOccupied   bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}
