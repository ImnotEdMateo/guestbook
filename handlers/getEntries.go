package handlers

import (
	"net/http"
	"github.com/ImnotEdMateo/guestbook/db"
	"github.com/ImnotEdMateo/guestbook/utils"
)

func GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	var entries []db.Entry
	if result := db.DB.Order("created_at DESC").Find(&entries); result.Error != nil {
		utils.ServeError(w, http.StatusInternalServerError, "Error al obtener las entradas")
		return
	}
	utils.ServeJSON(w, http.StatusOK, entries)
}
