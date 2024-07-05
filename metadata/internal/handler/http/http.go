package http

import (
	"errors"
	"log"
	"net/http"

	jason "github.com/micahasowata/jason/v2"
	"github.com/micahasowata/micro/metadata/internal/controller/metadata"
	"github.com/micahasowata/micro/metadata/internal/repository"
)

// Handler defines a movie metadata HTTP handler
type Handler struct {
	*jason.Parser
	ctrl *metadata.Controller
}

// New creates a new movie metadata HTTP handler
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{Parser: jason.Default(), ctrl: ctrl}
}

// GetMetadata handles GET /metadata requests
func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	if id == "" {
		h.Write(w, http.StatusBadRequest, "id not found")
		return
	}

	m, err := h.ctrl.Get(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			h.Write(w, http.StatusNotFound, err.Error())
			return
		}

		log.Printf("repository get error: %v\n", err)
		h.Write(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.Write(w, http.StatusOK, m)
}
