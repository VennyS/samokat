package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func SendMessage(w http.ResponseWriter, r *http.Request, message string, statusCode int) {
	render.Status(r, statusCode)
	render.JSON(w, r, map[string]string{"message": message})
}

func SendJSON(w http.ResponseWriter, r *http.Request, data map[string]any, statusCode int) {
	render.Status(r, statusCode)
	render.JSON(w, r, data)
}
