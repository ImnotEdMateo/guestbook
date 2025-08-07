package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/ImnotEdMateo/guestbook/db"
	"github.com/gorilla/mux"
)

// Utilidades
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, msg, status)
}

// Handlers
func GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	var entries []db.Entry
	if result := db.DB.Order("created_at DESC").Find(&entries); result.Error != nil {
		writeError(w, http.StatusInternalServerError, "Error al obtener las entradas")
		return
	}
	writeJSON(w, http.StatusOK, entries)
}

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
		writeError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var entry db.Entry
	if err := db.DB.First(&entry, id).Error; err != nil {
		writeError(w, http.StatusNotFound, "Entrada no encontrada")
		return
	}
	writeJSON(w, http.StatusOK, entry)
}

func handleRangeEntries(w http.ResponseWriter, rangeStr string) {
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		writeError(w, http.StatusBadRequest, "Formato de rango inválido (esperado: start-end)")
		return
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil || start < end {
		writeError(w, http.StatusBadRequest, "Rango inválido (start debe ser >= end)")
		return
	}

	var entries []db.Entry
	if err := db.DB.Where("id <= ? AND id >= ?", start, end).Order("id DESC").Find(&entries).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "Error al obtener las entradas")
		return
	}

	writeJSON(w, http.StatusOK, entries)
}

func PostEntryHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var entry db.Entry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		writeError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	if err := db.DB.Create(&entry).Error; err != nil {
		writeError(w, http.StatusInternalServerError, "Error guardando entrada: "+err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, entry)
}
