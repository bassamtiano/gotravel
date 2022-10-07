package tags_type_controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

var DB *sql.DB

func DBInit(Conn *sql.DB) {
	DB = Conn
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
