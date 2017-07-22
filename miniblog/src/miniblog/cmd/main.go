package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"miniblog"
	"miniblog/store"
	"time"
	"miniblog/entries"
)

func main() {
	// initialize store.
	var entryStore = store.NewMemoryStore()
	entryStore.Create(miniblog.Entry{
		Slug:       "go",
		Name:       "Go!",
		Body:       "Go Go Go!",
		CreateDate: time.Now(),
	})

	// initialize handler.
	var entryHandler = entries.NewHandler(entryStore)

	// create router and add routes.
	var router = mux.NewRouter()
	router.HandleFunc("/entries", entryHandler.GetAll).Methods("GET")
	router.HandleFunc("/entries/{pk}", entryHandler.Get).Methods("GET")
	http.ListenAndServe(":8080", router)
}
