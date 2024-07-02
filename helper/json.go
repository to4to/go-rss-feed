package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

// responseWithJSON writes a JSON response with the specified HTTP status code and payload to the provided http.ResponseWriter.
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed To Marshal Json Response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
