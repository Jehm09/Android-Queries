package main

import (
	"encoding/json"
	"log"
	"net/http"

	repo "./repository"
	"github.com/Jehm09/Android-Queries/server/model"
	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

// Create router
func New() Server {
	a := &api{}

	r := mux.NewRouter()

	//Domain methods get

	// r.HandleFunc("/domain", a.getDomains).Methods(http.MethodGet)
	r.HandleFunc("/domain/{value}", a.getDomain).Methods(http.MethodGet)

	//Histroy methods get

	// r.HandleFunc("/history", a.getHistory).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

// func (a *api) getDomains(w http.ResponseWriter, r *http.Request) {
// 	gophers, _ := a.repository.FetchGophers()

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(gophers)
// }

func (a *api) getDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	URLID := vars["value"]
	history := model.History{Items: make([]string, 0, 100)}

	domain := repo.GetDomain(URLID, &history)
	w.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
	// 	json.NewEncoder(w).Encode("Gopher Not found")
	// 	return
	// }

	json.NewEncoder(w).Encode(domain)
}

//Main
func main() {
	s := New()

	log.Fatal(http.ListenAndServe(":8070", s.Router()))
}
