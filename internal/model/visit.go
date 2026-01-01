package model

import (
	"time"

	"github.com/google/uuid"
)

type Visit struct {
	ID            uuid.UUID     `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProfileID     uuid.UUID     `gorm:"type:uuid;not null;index"`
	Profile       Profile       `gorm:"foreignKey:ProfileID;references:ID"`
	ArrivalDate   time.Time     `gorm:"not null"`
	DepartureDate *time.Time    `gorm:"default:null"`
	StayAreaID    uuid.UUID     `gorm:"type:uuid;not null;index"`
	StayArea      StayArea      `gorm:"foreignKey:StayAreaID;references:ID"`
	Status        ProfileStatus `gorm:"type:varchar(20);not null;default:'pending'"`
	LockerID      *uuid.UUID    `gorm:"type:uuid;index"`
	Locker        *Locker       `gorm:"foreignKey:LockerID;references:ID"`
	Remarks       *string       `gorm:"type:text"`
	CreatedAt     time.Time     `gorm:"autoCreateTime"`
}

type ProfileStatus string

const (
	StatusCheckedIn  ProfileStatus = "checked-in"
	StatusPending    ProfileStatus = "pending"
	StatusCheckedOut ProfileStatus = "checked-out"
)
