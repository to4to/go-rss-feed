package handler

import (
	"encoding/json"
	"log"
	"net/http"
)



func responseWithJson(w http.ResponseWriter, /* Status code*/code int,payload interface{}){

data,err:=json.Marshal(payload)

if err!=nil{
	log.Printf("Failed To MArshal Json: %v",payload)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal Server Error"))
	return
}


w.Header().Set("Content-Type","application/json")
w.WriteHeader(code)
w.Write(data)
}