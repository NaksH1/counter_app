package model

import (
	"time"

	"github.com/google/uuid"
)

type Feedback struct {
	ID        uuid.UUID    `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	ProfileID uuid.UUID    `gorm:"type:uuid;not null;index"`
	Profile   Profile      `gorm:"foreignKey:ProfileID;references:ID"`
	VisitID   *uuid.UUID   `gorm:"type:uuid;index"`
	Visit     *Visit       `gorm:"foreignKey:VisitID;references:ID"`
	Content   string       `gorm:"type:text;not null"`
	Type      FeedbackType `gorm:"type:varchar(20);not null"`
	CreatedBy *string      `gorm:"type:text"`
	CreatedAt time.Time    `gorm:"autoCreateTime"`
}

type FeedbackType string

const (
	TypePositive FeedbackType = "Positive"
	TypeNegative FeedbackType = "Negative"
	TypeNeutral  FeedbackType = "Neutral"
)
