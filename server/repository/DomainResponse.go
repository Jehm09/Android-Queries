package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jehm09/Android-Queries/server/model"
	"github.com/badoux/goscraper"
)

//This API gives us information about the domain, its servers and the ssl grade of each server
const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="
const PREFIX_URL = "://www."

// type ActualData struct {
// 	History *model.History
// 	Domain  *model.Domain
// }

// func NewActualData(history *model.History, domain *model.Domain) *ActualData {
// 	return &ActualData{History: history, Domain: domain}
// }

func GetDomain(host string, history *model.History) *model.Domain {
	response, err := http.Get(API_DOMAINS_URL + host)

	if err != nil {
		log.Fatal(err.Error())
		// os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var domainA DomainAPI
	json.Unmarshal([]byte(responseData), &domainA)

	return createDomain(domainA, history)
}

func createDomain(domainA DomainAPI, history *model.History) *model.Domain {
	// result := &models.Domain{}

	// Agrego al historial si no existe la busqueda
	if !history.Exist(domainA.Host) {
		history.Items = append(history.Items, domainA.Host)
	}

	// Metodo que consulte si existe en la base de datos

	// si no existe se crea el primero

	// Creo un dominio
	domainResults := model.Domain{Servers: make([]model.Server, 0, 100)}

	// Si el servidor esta caido
	if len(domainA.Erros) > 0 {
		domainResults.IsDown = true
	} else {
		FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
		domainResults.ServersChanged = false
		domainResults.SslGrade = domainA.SearchMinorGrade()
		domainResults.PreviousSslGrade = ""
		domainResults.IsDown = false
		domainResults.Title, domainResults.Logo = getPageInfo(FullUrl)
		// timeActual := time.Now()
		createServersOfDomain(domainA, &domainResults)
		// domainResults.Time = time.Now()
	}

	return &domainResults
}

//Crea los servidores
func createServersOfDomain(domainA DomainAPI, domain *model.Domain) {
	owner, country := domainA.WhoisServerAttributes()

	for _, servers := range domainA.Endpoints {
		address := servers.IPAddress
		sslGrade := servers.Grade
		temServer := model.Server{address, sslGrade, country, owner}
		domain.Servers = append(domain.Servers, temServer)
	}
}

// Obteiene el titulo y el logo de la pagina
func getPageInfo(url string) (string, string) {
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		log.Println(err)
		return "", ""
	}
	return s.Preview.Title, s.Preview.Icon
}
