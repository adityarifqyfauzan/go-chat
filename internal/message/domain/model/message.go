package model

import "time"

type Message struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `gorm:"index"`
	RoomID    uint `gorm:"index"`
	Content   string
	CreatedAt time.Time
}
