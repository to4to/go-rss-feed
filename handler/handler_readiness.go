package handler

import "net/http"



func HandlerReadiness(w http.ResponseWriter,r *http.Request){

	responseWithJson(w,200,struct{}{})
}