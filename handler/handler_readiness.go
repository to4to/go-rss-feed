package handler

import "net/http"



func handlerReadiness(w http.ResponseWriter,r *http.Request){

	responseWithJson(w,200,struct{}{})
}