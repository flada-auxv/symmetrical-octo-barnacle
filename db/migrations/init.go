package main

import (
	"log"

	"github.com/flada-auxv/symmetrical-octo-barnacle/domain/entity"
	"github.com/flada-auxv/symmetrical-octo-barnacle/infra/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal("failed to connect database")
	}

	defer db.Close()

	db.AutoMigrate(
		&entity.User{},
		&entity.Task{},
	)
}
