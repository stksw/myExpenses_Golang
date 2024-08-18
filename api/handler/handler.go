package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) RenderResponse(w http.ResponseWriter, r *http.Request, data any) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]any{})
}

func (h Handler) RenderErrorResponse(w http.ResponseWriter, r *http.Request) {

}
