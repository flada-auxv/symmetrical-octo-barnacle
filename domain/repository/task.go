package repository

import "github.com/flada-auxv/symmetrical-octo-barnacle/domain/entity"

type TaskRepository interface {
	FindAll() ([]*entity.Task, error)
}
