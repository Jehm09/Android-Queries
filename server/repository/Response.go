package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

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

	return CreateDomain(domainA, history)
}

func CreateDomain(domainA DomainAPI, history *model.History) *model.Domain {
	// result := &models.Domain{}
	if history != nil {
		history.Items = append(history.Items, domainA.Host)
	} else {
		// var items []string
		// history = &models.History{items}
	}

	// Metodo que consulte si existe en la base de datos

	// si no existe se crea el primero
	var domainResults model.Domain

	// Si el servidor esta caido
	if len(domainA.Erros) > 0 {
		domainResults.IsDown = true
	} else {
		FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
		domainResults.ServersChanged = false
		domainResults.SslGrade = domainA.SearchMinorGrade()
		domainResults.PreviousSslGrade = ""
		domainResults.IsDown = false
		domainResults.Title, domainResults.Logo = GetPageInfo(FullUrl)
		timeActual := time.Now()
		// domainResults.Time = time.Now()
	}

}

func GetPageInfo(url string) (string, string) {
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		log.Println(err)
		return "", ""
	}
	return s.Preview.Title, s.Preview.Icon
}
