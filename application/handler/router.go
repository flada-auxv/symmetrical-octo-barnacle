package handler

import (
	"log"

	"github.com/flada-auxv/symmetrical-octo-barnacle/infra/db"
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
