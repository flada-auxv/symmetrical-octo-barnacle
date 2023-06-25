package handler

import (
	"net/http"
	"time"

	"github.com/flada-auxv/symmetrical-octo-barnacle/application/usecase"
	"github.com/labstack/echo/v4"
)

func GetTasks(c echo.Context) error {
	uc := usecase.NewGetTasks(c.(*AppContext).taskRepo)
	tasks, err := uc.Execute()
	if err != nil {
		// TODO: error response
		return err
	}

	res := make(GetTasksResponse, len(tasks))

	for _, t := range tasks {
		res = append(res, GetTasksResponseItem{
			UserID:      t.UserID,
			Title:       t.Title,
			Description: t.Description,
			StartedAt:   t.StartedAt,
			CompletedAt: t.CompletedAt,
		})
	}

	return c.JSON(http.StatusOK, res)
}

type GetTasksResponse []GetTasksResponseItem
type GetTasksResponseItem struct {
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
}
