package model

import "time"

type Room struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
}
