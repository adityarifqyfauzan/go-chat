package repository

import (
	"github.com/adityarifqyfauzan/go-chat/internal/authentication/domain/model"
	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string, page, size int) ([]*model.User, error)
	CountByUsername(username string) int64
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindByUsername(username string, page, size int) ([]*model.User, error) {
	m := make([]*model.User, 0)
	limit, offset := utils.GetLimitOffset(page, size)

	if err := r.db.Where("username LIKE ?", "%"+username+"%").Limit(limit).Offset(offset).Find(&m).Error; err != nil {
		return nil, err
	}

	return m, nil
}

func (r *userRepository) CountByUsername(username string) int64 {
	var count int64
	r.db.Model(&model.User{}).Where("username LIKE ?", "%"+username+"%").Count(&count)
	return count
}
