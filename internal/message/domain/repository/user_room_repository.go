package repository

import (
	"github.com/adityarifqyfauzan/go-chat/internal/message/domain/model"
	"gorm.io/gorm"
)

type UserRoomRepository interface {
	FindOneBy(criteria map[string]interface{}) (*model.UserRoom, error)
	Create(tx *gorm.DB, m *model.UserRoom) error
	CheckExistingRoom(senderID, receiverID uint) (*model.Room, error)
	DeleteUserRoom(userRoomIDs ...uint) error
}

type userRoomRepository struct {
	db *gorm.DB
}

func NewUserRoomRepository(db *gorm.DB) UserRoomRepository {
	return &userRoomRepository{db: db}
}

func (urr *userRoomRepository) FindOneBy(criteria map[string]interface{}) (*model.UserRoom, error) {
	m := new(model.UserRoom)
	if err := urr.db.Where(criteria).First(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (urr *userRoomRepository) Create(tx *gorm.DB, m *model.UserRoom) error {
	return tx.Create(m).Error
}

func (urr *userRoomRepository) CheckExistingRoom(senderID, receiverID uint) (*model.Room, error) {
	var m *model.Room
	if err := urr.db.Raw(`
		SELECT r.*
			FROM user_rooms ur1
			JOIN user_rooms ur2 ON ur1.room_id = ur2.room_id
			JOIN rooms r ON ur1.room_id = r.id
			WHERE ur1.user_id = ? AND ur2.user_id = ?;
	`, senderID, receiverID).Scan(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (urr *userRoomRepository) DeleteUserRoom(userRoomIDs ...uint) error {
	return urr.db.Delete(&model.UserRoom{}, userRoomIDs).Error
}
