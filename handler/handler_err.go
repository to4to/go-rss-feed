package handler

import "net/http"


func HandlerErr(w http.ResponseWriter,r *http.Request) {
	respondWithError(w,200,"Something Went Wrong")
}