package repository

import (
	"database/sql"
	"log"

	database "github.com/Jehm09/Android-Queries/server/database"
	"github.com/Jehm09/Android-Queries/server/model"
)

func GetHistory(db *sql.DB) *model.History {
	historyRepo := database.NewHistoyRepository(db)
	items, err := historyRepo.FetchHistory()

	if err != nil {
		log.Fatal(err)
	}

	return &model.History{Items: items}
}
