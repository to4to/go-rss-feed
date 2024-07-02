package helper

import (
	"encoding/json"
	"log"
	"net/http"
)


// RespondWithJSON writes a JSON response to the http.ResponseWriter with the specified status code and payload.
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    // Marshal the payload into JSON format
    data, err := json.Marshal(payload)

    // Check for any errors during marshaling
    if err != nil {
        log.Printf("Failed To Marshal Json Response: %v", payload)
        w.WriteHeader(500)
        return
    }

    // Set the Content-Type header to application/json
    w.Header().Add("Content-Type", "application/json")
    // Set the status code of the response
    w.WriteHeader(code)
    // Write the JSON data to the response writer
    w.Write(data)
}
