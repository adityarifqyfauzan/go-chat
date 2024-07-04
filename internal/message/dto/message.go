package dto

import (
	"time"

	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/model"
)

type RoomCreateRequest struct {
	AuthenticatedUser uint
	FriendID          uint `json:"friend_id"`
}

type RoomRequest struct {
	AuthenticatedUser uint
	Page              int `form:"page"`
	Size              int `form:"size"`
}

type ConversationRequest struct {
	AuthenticatedUser uint
	RoomID            uint `uri:"room_id"`
	Page              int  `form:"page"`
	Size              int  `form:"size"`
}

type RoomResponse struct {
	ID        int       `json:"id"`
	RoomID    int       `json:"room_id"`
	UserID    int       `json:"user_id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (rr *RoomResponse) ModelToDto(m *model.UserRoomDetail) {
	rr.ID = int(m.ID)
	rr.RoomID = int(m.RoomID)
	rr.UserID = int(m.UserID)
	rr.Username = m.Username
	rr.Name = m.Name
}

type ConversationResponse struct {
	ID                int       `json:"id"`
	RoomID            int       `json:"room_id"`
	UserID            int       `json:"user_id"`
	Role              string    `json:"as"`
	Content           string    `json:"content"`
	CreatedAt         time.Time `json:"created_at"`
	AuthenticatedUser uint      `json:"-"`
}

func (cr *ConversationResponse) ModelToDto(m *model.Message) {
	cr.ID = int(m.ID)
	cr.RoomID = int(m.RoomID)
	cr.UserID = int(m.UserID)
	cr.Role = "receiver"
	if cr.AuthenticatedUser == m.UserID {
		cr.Role = "sender"
	}
	cr.Content = m.Content
	cr.CreatedAt = m.CreatedAt
}
