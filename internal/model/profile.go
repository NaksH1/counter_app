package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name        string    `gorm:"not null"`
	Email       string    `gorm:"unique;not null"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Gender      Gender    `gorm:"type:varchar(10);not null"`
	Category    Category  `gorm:"type:varchar(300)"`
	IsBlocked   bool      `gorm:"default:false"`
	Remarks     *string   `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

type Category string

const (
	CategorySTV      Category = "Short Term Volunteer"
	CategoryLTV      Category = "Long Term Volunteer"
	CategoryOverseas Category = "Overseas Volunteer"
)
