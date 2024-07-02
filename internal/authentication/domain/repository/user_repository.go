package repository

import (
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindOneBy(criteria map[string]interface{}) (*model.User, error)
	Create(*model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindOneBy(criteria map[string]interface{}) (*model.User, error) {
	var m *model.User
	if err := r.db.Where(criteria).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *userRepository) Create(m *model.User) error {
	return r.db.Create(m).Error
}
