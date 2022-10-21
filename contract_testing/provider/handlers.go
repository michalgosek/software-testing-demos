package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type MessageWithIDRepository interface {
	MessageWithID(ID string) Message
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "\t")
		err := enc.Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func MessageWithIDHandler(c MessageWithIDRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID := chi.URLParam(r, "ID")
		if ID == "" {
			SendJSONResponse(w, "{ message: missing ID value in the endpoint. }", http.StatusBadRequest)
			return
		}
		m := c.MessageWithID(ID)
		if m.IsEmpty() {
			SendJSONResponse(w, m, http.StatusNotFound)
		}
		SendJSONResponse(w, &m, http.StatusOK)
	}
}

type MessagesRepository interface {
	Messages() []Message
}

func MessagesHandler(c MessagesRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mm := c.Messages()
		if len(mm) == 0 {
			SendJSONResponse(w, mm, http.StatusNotFound)
		}
		SendJSONResponse(w, mm, http.StatusOK)
	}
}

func NewMux(cache *MessageCache) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api/v1/messages", func(r chi.Router) {
		r.Get("/", MessagesHandler(cache))
		r.Get("/{ID}", MessageWithIDHandler(cache))
	})
	return r
}
