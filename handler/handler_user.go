package handler

import "net/http"



func HandlerCreateUser(w http.ResponseWriter,r *http.Request){

	responseWithJson(w,200,struct{}{})
}