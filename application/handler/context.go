package handler

import (
	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/repository"
	db_repository "github.com/flada-auxv/symmetrical-octo-barnacle/infra/repository/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	echo.Context
	RepositoryContext
}

type RepositoryContext struct {
	taskRepo repository.TaskRepository
}

func NewAppContext(c echo.Context, db *gorm.DB) *AppContext {
	return &AppContext{
		c,
		RepositoryContext{
			taskRepo: db_repository.NewDBTaskRepositry(db),
		},
	}
}
