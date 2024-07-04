package model

import "time"

type UserRoom struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `gorm:"index"`
	RoomID    uint `gorm:"index"`
	CreatedAt time.Time
}

type UserRoomDetail struct {
	UserRoom
	Username string
	Name     string
}
