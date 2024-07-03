package handler

import (
	"net/http"

	"github.com/to4to/go-rss-feed/helper"
	"github.com/to4to/go-rss-feed/models"
)


type LocalApiConfig struct{
	models.ApiConfig
}
func (l *LocalApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {

	helper.RespondWithJSON(w ,200,struct{}{})
}
