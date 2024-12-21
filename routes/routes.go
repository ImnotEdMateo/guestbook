package routes

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/ImnotEdMateo/guestbook/db"
  "github.com/gorilla/mux"
)

func GetEntriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var entries []db.Entry
	if result := db.DB.Order("created_at DESC").Find(&entries); result.Error != nil {
		http.Error(w, "Error al obtener las entradas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func GetEntryHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, err := strconv.Atoi(vars["id"])
  if err != nil {
  	w.WriteHeader(http.StatusBadRequest)
  	w.Write([]byte("Invalid entry ID"))
  	return
  }
  
  var entry db.Entry
  if err := db.DB.First(&entry, id).Error; err != nil {
  	w.WriteHeader(http.StatusNotFound)
  	w.Write([]byte("Entry not found"))
  	return
  }
  
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&entry)
}

func PostEntryHandler(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  
  var entry db.Entry
  if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
  	w.WriteHeader(http.StatusBadRequest)
  	w.Write([]byte("Invalid request payload"))
  	return
  }
  
  if err := db.DB.Create(&entry).Error; err != nil {
  	w.WriteHeader(http.StatusInternalServerError)
  	w.Write([]byte("Error saving entry: " + err.Error()))
  	return
  }
  
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(&entry)
}
