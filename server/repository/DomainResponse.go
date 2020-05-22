package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	database "github.com/Jehm09/Android-Queries/server/database"
	"github.com/Jehm09/Android-Queries/server/model"
	"github.com/badoux/goscraper"
)

//This API gives us information about the domain, its servers and the ssl grade of each server
const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="
const PREFIX_URL = "://www."
const DEFAUlT_GRADE = "-"

// type ActualData struct {
// 	History *model.History
// 	Domain  *model.Domain
// }

// func NewActualData(history *model.History, domain *model.Domain) *ActualData {
// 	return &ActualData{History: history, Domain: domain}
// }

func GetDomain(host string, db *sql.DB) *model.Domain {
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

	return createDomain(domainA, db)
}

func createDomain(domainA DomainAPI, db *sql.DB) *model.Domain {
	// result := &models.Domain{}

	// Se crean las variables para poder consultar la base de datos
	historyDb := database.NewHistoyRepository(db)
	domainDb := database.NewDomainRepository(db)

	// Agrego al historial el hostname
	err := historyDb.CreateHistory(domainA.Host)

	// Ocurrio un error al agregar el hostname al historial
	if err != nil {
		fmt.Println(err)
	}

	// Metodo que consulte si existe en la base de datos
	domainExists, err := domainDb.FetchDomain(domainA.Host)
	if err != nil {
		fmt.Println(err)
	}

	// Creo un dominio
	domainResults := model.Domain{Servers: make([]model.Server, 0, 100)}

	// El dominio ya estaba en la base de datos
	if domainExists != nil {
		// Si el servidor esta caido
		if len(domainA.Erros) > 0 {
			domainResults.IsDown = true
		} else {
			FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
			// Si paso una hora
			domainResults.SslGrade = domainA.SearchMinorGrade()
			domainResults.ServersChanged = domainResults.SslGrade != domainExists.SslGrade
			// Si paso una hora
			domainResults.PreviousSslGrade = domainExists.SslGrade
			domainResults.IsDown = false
			domainResults.Title, domainResults.Logo = getPageInfo(FullUrl)
			createServersOfDomain(domainA, &domainResults)

			//actualizo en la base de datos
			domainData := database.DomainDB{domainA.Host, domainResults.SslGrade, domainResults.PreviousSslGrade, time.Now()}
			// Error si no guarda
			err := domainDb.UpdateDomain(&domainData)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		// No existia en la base de datos
		if len(domainA.Erros) > 0 {
			domainResults.IsDown = true
		} else {
			FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
			domainResults.ServersChanged = false
			domainResults.SslGrade = domainA.SearchMinorGrade()
			domainResults.PreviousSslGrade = "-"
			domainResults.IsDown = false
			domainResults.Title, domainResults.Logo = getPageInfo(FullUrl)

			//Guardo en la base de datos
			domainData := database.DomainDB{domainA.Host, domainResults.SslGrade, "-", time.Now()}
			// Error si no guarda
			err := domainDb.CreateDomain(&domainData)
			if err != nil {
				log.Fatal(err)
			}
			// Creo servidores
			createServersOfDomain(domainA, &domainResults)
		}
	}

	return &domainResults
}

// func updateDomainDB(domainA DomainAPI, domainDb *database.domainRepo) {

// }

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

// Mira si ya paso una hora
func CompareOneHourBefore(lastSearch time.Time) bool {
	today := time.Now()
	comparator := lastSearch.Add(1 * time.Hour)
	if today.Before(comparator) {
		return false
	} else {
		return true
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
