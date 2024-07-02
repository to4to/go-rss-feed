package handler

import (
	"net/http"

	"github.com/to4to/go-rss-feed/helper"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {

	helper.RespondWithJSON(w, 200, struct{}{})

}
