package entries

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"errors"
	"miniblog/store"
)

type Handler struct {
	store store.Store
}

func NewHandler(courseStore store.Store) *Handler {
	return &Handler{
		store: courseStore,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.store.GetAll())
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	var vars = mux.Vars(r)

	slug, ok := vars["pk"]
	if !ok {
		panic(errors.New("Whoa whoa whoa, invalid course pk!"))
	}

	course, err := h.store.Get(slug)
	if err != nil {
		panic(errors.New("Whoa whoa whoa, that course does not exists!"))
	}

	json.NewEncoder(w).Encode(course)
}
