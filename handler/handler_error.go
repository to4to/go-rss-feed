package handler

import (
	"net/http"

	"github.com/to4to/go-rss-feed/helper"
)

// HandlerError is a function that handles errors by responding with a specific status code and error message.
// It takes in the http.ResponseWriter interface to write the response back to the client and the http.Request object for additional context.
func HandlerError(w http.ResponseWriter, r *http.Request) {
    helper.RespondWithError(w, 400, "Something went Wrong")
}
