package model

import (
	"github.com/google/uuid"
	"time"
)

type Item struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Category *Category `gorm:"foreignkey:Category",json:"category"`
	AddedAt  time.Time `json:"addedAt"`
	Quantity uint      `json:"quantity"`
}
