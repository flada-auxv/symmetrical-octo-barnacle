package db_repository

import (
	"github.com/jinzhu/gorm"

	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/entity"
)

type DBTaskRepository struct {
	db *gorm.DB
}

func NewDBTaskRepositry(db *gorm.DB) *DBTaskRepository {
	return &DBTaskRepository{
		db: db,
	}
}

func (r *DBTaskRepository) FindAll() ([]*entity.Task, error) {
	var tasks []*entity.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
