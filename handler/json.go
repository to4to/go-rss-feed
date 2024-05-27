package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithJson(w http.ResponseWriter /* Status code*/, code int, payload interface{}) {

	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed To MArshal Json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}


func respondWithError(w http.ResponseWriter, code int, message string) {


	//error code in 400s are client side error
if code>499{
log.Println("Responding With 5XX Error:" ,message)
}



type errorResponse struct {
	Error string `json:"error"`
}
	responseWithJson(w, code, errorResponse{Error: message})
}