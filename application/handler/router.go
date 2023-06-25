package handler

import (
	"log"

	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/repository"
	"github.com/flada-auxv/symmetrical-octo-barnacle/infra/db"
	db_repository "github.com/flada-auxv/symmetrical-octo-barnacle/infra/repository/db"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	db, err := db.Connect()
	if err != nil {
		log.Fatal("failed to connect database")
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(NewAppContext(c, db))
		}
	})

	e.GET("/tasks", GetTasks)

	return e
}

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
