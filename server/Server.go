package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	conec "./database"
	repo "./repository"
	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
	db     *sql.DB
}

type Server interface {
	Router() http.Handler
}

// Create router
func New() Server {
	a := &api{}

	r := mux.NewRouter()

	//Domain methods get
	r.HandleFunc("/domain/{value}", a.getDomain).Methods(http.MethodGet)

	// Histroy methods get
	r.HandleFunc("/history", a.getHistory).Methods(http.MethodGet)

	// r.HandleFunc("/history", a.getHistory).Methods(http.MethodGet)

	db, err := conec.GetConnectionDB()
	if err != nil {
		log.Fatal(err)
	}

	a.router = r
	a.db = db
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) getDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	URLID := vars["value"]

	domain := repo.GetDomain(URLID, a.db)
	// w.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
	// 	json.NewEncoder(w).Encode("Gopher Not found")
	// 	return
	// }

	json.NewEncoder(w).Encode(domain)
}

func (a *api) getHistory(w http.ResponseWriter, r *http.Request) {
	history := repo.GetHistory(a.db)

	// w.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound) // We use not found for simplicity
	// 	json.NewEncoder(w).Encode("Gopher Not found")
	// 	return
	// }

	json.NewEncoder(w).Encode(history)
}

//Main
func main() {
	s := New()

	log.Fatal(http.ListenAndServe(":8070", s.Router()))
}
