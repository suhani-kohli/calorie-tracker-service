package api

import (
	"net/http"

	"gorm.io/gorm"
)

func GetItem(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
