package repository

import (
	"time"

	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/model"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"gorm.io/gorm"
)

type RoomRepository interface {
	FindByUserID(userID uint, page, size int) ([]*model.UserRoomDetail, error)
	Count(userID uint) int64
	Create(tx *gorm.DB) (uint, error)
	FindRoomBy(criteria map[string]interface{}) ([]*model.Room, error)
	DeleteRoom(roomID uint) error
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) FindByUserID(userID uint, page, size int) ([]*model.UserRoomDetail, error) {
	m := make([]*model.UserRoomDetail, 0)
	limit, offset := utils.GetLimitOffset(page, size)

	if err := r.db.Raw(`
			SELECT ur.*, u.username, u.name FROM user_rooms ur
			JOIN users u ON u.id = ur.user_id
			WHERE ur.user_id != ? AND ur.room_id IN 
				(SELECT id FROM rooms WHERE id IN (SELECT room_id FROM user_rooms WHERE user_id = ?)) LIMIT ? OFFSET ?
		`, userID, userID, limit, offset).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *roomRepository) Count(userID uint) int64 {
	var count int64
	r.db.Raw(`
			WITH latest_messages AS (
				SELECT 
					m.id as message_id,
					m.room_id,
					m.created_at,
					ROW_NUMBER() OVER (PARTITION BY m.room_id ORDER BY m.created_at DESC) as row_num
				FROM messages m
			)
			SELECT ur.id
			FROM users u JOIN user_rooms ur ON ur.user_id = u.id
			JOIN rooms r ON r.id = ur.room_id
			JOIN latest_messages lm ON lm.room_id = r.id
			WHERE ur.user_id != ? AND lm.row_num = 1
			ORDER BY lm.created_at DESC
		`, int(userID)).Count(&count)
	return count
}

func (r *roomRepository) Create(tx *gorm.DB) (uint, error) {
	m := model.Room{
		CreatedAt: time.Now(),
	}
	if err := r.db.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *roomRepository) FindRoomBy(criteria map[string]interface{}) ([]*model.Room, error) {
	m := make([]*model.Room, 0)
	if err := r.db.Where(criteria).Find(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (r *roomRepository) DeleteRoom(roomID uint) error {
	return r.db.Delete(&model.Room{}, roomID).Error
}
