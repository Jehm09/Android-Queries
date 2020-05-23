package main

import (
	"database/sql"
	"encoding/json"
	"log"

	conec "./database"
	repo "./repository"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

const ADDRESS = "192.168.0.15"
const PORT = "8070"

type api struct {
	router *fasthttprouter.Router
	db     *sql.DB
}

type Server interface {
	Router() *fasthttprouter.Router
}

// Create router
func New() Server {
	a := &api{}

	r := fasthttprouter.New()

	//Domain methods get
	r.GET("/domain", a.getDomain)

	// Histroy methods get
	r.GET("/history", a.getHistory)

	db, err := conec.GetConnectionDB()
	if err != nil {
		log.Fatal(err)
	}

	a.router = r
	a.db = db

	return a
}

func (a *api) Router() *fasthttprouter.Router {
	return a.router
}

// localhost:8070/domain?hostname=x
func (a *api) getDomain(ctx *fasthttp.RequestCtx) {
	nameByte := ctx.QueryArgs().Peek("hostname")

	hostName := string(nameByte)

	domain := repo.GetDomain(hostName, a.db)

	jsonBody, err := json.Marshal(domain)

	if err != nil {
		ctx.Error(" json marshal fail", 500)
		return
	}

	ctx.Response.SetBody(jsonBody)

	// json.NewEncoder(w).Encode(domain)
}

func (a *api) getHistory(ctx *fasthttp.RequestCtx) {
	history := repo.GetHistory(a.db)

	jsonBody, err := json.Marshal(history)

	if err != nil {
		ctx.Error(" json marshal fail", 500)
		return
	}

	ctx.Response.SetBody(jsonBody)
}

//Main
func main() {
	s := New()

	serverListener := (ADDRESS + ":" + PORT)
	log.Fatal(fasthttp.ListenAndServe(serverListener, s.Router().Handler))
}
