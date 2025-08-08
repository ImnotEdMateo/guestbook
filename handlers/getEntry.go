package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ImnotEdMateo/guestbook/utils"
	"github.com/ImnotEdMateo/guestbook/db"
)

func GetEntryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	
	// Verifica si es un rango tipo 59-10
	if strings.Contains(idStr, "-") {
		handleRangeEntries(w, idStr)
		return
	}

	// Entrada individual
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ServeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var entry db.Entry
	if err := db.DB.First(&entry, id).Error; err != nil {
		utils.ServeError(w, http.StatusNotFound, "Entrada no encontrada")
		return
	}
	utils.ServeJSON(w, http.StatusOK, entry)
}

func handleRangeEntries(w http.ResponseWriter, rangeStr string) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		utils.ServeError(w, http.StatusBadRequest, "Formato de rango inválido (esperado: start-end)")
		return
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || start < end {
		utils.ServeError(w, http.StatusBadRequest, "Rango inválido (start debe ser >= end)")
		return
	}

	var entries []db.Entry
	if err := db.DB.Where("id <= ? AND id >= ?", start, end).Order("id DESC").Find(&entries).Error; err != nil {
		utils.ServeError(w, http.StatusInternalServerError, "Error al obtener las entradas")
		return
	}

	utils.ServeJSON(w, http.StatusOK, entries)
}

func PostEntryHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var entry db.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		utils.ServeError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	if err := db.DB.Create(&entry).Error; err != nil {
		utils.ServeError(w, http.StatusInternalServerError, "Error guardando entrada: "+err.Error())
		return
	}

	utils.ServeJSON(w, http.StatusCreated, entry)
}
