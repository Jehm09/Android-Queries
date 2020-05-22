package repository

import (
	"database/sql"

	database "github.com/Jehm09/Android-Queries/server/database"
	"github.com/Jehm09/Android-Queries/server/model"
)

func GetHistory(db *sql.DB) *model.History {
	historyRepo := database.NewHistoyRepository(db)
	items, err := historyRepo.FetchHistory()

	return &model.History{Items: items}
}
