package usecase

import (
	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/entity"
	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/repository"
)

type GetTasks struct {
	repo repository.TaskRepository
}

func NewGetTasks(repo repository.TaskRepository) *GetTasks {
	return &GetTasks{repo: repo}
}

func (uc GetTasks) Execute() ([]*entity.Task, error) {
	return uc.repo.FindAll()
}
