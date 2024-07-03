package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/to4to/go-rss-feed/helper"
	"github.com/to4to/go-rss-feed/internal/db"
	"github.com/to4to/go-rss-feed/models"
)

type LocalApiConfig struct {
	models.ApiConfig
}

func (l *LocalApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramaters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}
	err := decoder.Decode(&params)

	if err != nil {
		helper.RespondWithError(w, 400, fmt.Sprintf("Error Parsing JSON", err))
		return
	}

	user,err:= l.DB.CreateUser(r.Context(),db.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now(),
		Name: params.Name,
	})


	if err!=nil{
		helper.RespondWithError(w, 400, fmt.Sprintf("Couldn't Create User", err))
		return
	}
	helper.RespondWithJSON(w, 200, user)
}
