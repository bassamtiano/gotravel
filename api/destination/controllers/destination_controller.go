package destination_controller

import (
	"encoding/json"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world" }`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

func AddDestinationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
}
