package repository

import (
	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/model"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"gorm.io/gorm"
)

type MessageRepository interface {
	FindConversation(roomID uint, page, size int) ([]*model.Message, error)
	CountConversation(roomID uint) int64
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) FindConversation(roomID uint, page, size int) ([]*model.Message, error) {
	m := make([]*model.Message, 0)
	limit, offset := utils.GetLimitOffset(page, size)
	if err := r.db.Where("room_id = ?", roomID).Order("id DESC").Limit(limit).Offset(offset).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (r *messageRepository) CountConversation(roomID uint) int64 {
	var count int64
	r.db.Where("room_id = ?", roomID).Count(&count)
	return count
}
