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
const DEFAULT_GRADE = "-"

func GetDomain(host string, db *sql.DB) *model.Domain {
	response, err := http.Get(API_DOMAINS_URL + host)

	if err != nil {
		log.Fatal(err.Error())
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

	historyDb := database.NewHistoyRepository(db)
	domainDb := database.NewDomainRepository(db)

	// Add hostname to history
	err := historyDb.CreateHistory(domainA.Host)

	// Error to put hostname in the database
	if err != nil {
		fmt.Println(err)
	}

	// Get the domain of the database with that host of consult
	domainExists, err := domainDb.FetchDomain(domainA.Host)
	if err != nil {
		fmt.Println(err)
	}

	// Created a domain that goin to return
	domainResults := model.Domain{Servers: make([]model.Server, 0, 100)}

	// The domain exists in the databas
	if domainExists != nil {
		// The api rest api.ssllabs.com can not get results
		if len(domainA.Erros) > 0 {
			domainResults = *newDomain(false, "", "", "", "", true)
		} else {
			FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
			title, logo := getPageInfo(FullUrl)

			// If it's been an hour
			sslGrade, serverChanged, previousSslGrade := DEFAULT_GRADE, false, DEFAULT_GRADE
			if CompareOneHourBefore(domainExists.LastSearch) {
				sslGrade = domainA.SearchMinorGrade()
				serverChanged = sslGrade != domainExists.SslGrade
				previousSslGrade = domainExists.SslGrade
			}

			domainResults = *newDomain(serverChanged, sslGrade, previousSslGrade, logo, title, false)
			createServersOfDomain(domainA, &domainResults)

			// Update the domain on the database
			domainData := database.DomainDB{domainA.Host, domainResults.SslGrade, domainResults.PreviousSslGrade, time.Now()}

			// Error, in case the database can't save
			err := domainDb.UpdateDomain(&domainData)
			if err != nil {
				log.Fatal(err)
			}
		}

	} else {
		// The host does exists in database
		if len(domainA.Erros) > 0 {
			domainResults = *newDomain(false, "", "", "", "", true)
		} else {
			FullUrl := domainA.Protocol + PREFIX_URL + domainA.Host
			title, logo := getPageInfo(FullUrl)
			minorGrade := domainA.SearchMinorGrade()
			domainResults = *newDomain(false, minorGrade, DEFAULT_GRADE, logo, title, false)

			// Create a variable domainData
			domainData := database.DomainDB{domainA.Host, domainResults.SslGrade, "-", time.Now()}
			// Error, in case the database can't save
			err := domainDb.CreateDomain(&domainData)
			if err != nil {
				log.Fatal(err)
			}
			// I add the servers to domainResults
			createServersOfDomain(domainA, &domainResults)
		}
	}

	// Return domainResults
	return &domainResults
}

func newDomain(serversChanged bool, sslGrade string, previousSsl string, logo string, title string, isDown bool) *model.Domain {
	return &model.Domain{ServersChanged: serversChanged, SslGrade: sslGrade, PreviousSslGrade: previousSsl, Title: title, Logo: logo, IsDown: isDown}
}

//Created the servers of a domain
func createServersOfDomain(domainA DomainAPI, domain *model.Domain) {

	for _, servers := range domainA.Endpoints {
		address := servers.IPAddress
		owner, country := domainA.WhoisServerAttributes(address)
		sslGrade := servers.Grade
		temServer := model.Server{address, sslGrade, country, owner}
		domain.Servers = append(domain.Servers, temServer)
	}
}

// If it's been an hour
func CompareOneHourBefore(lastSearch time.Time) bool {
	today := time.Now()
	comparator := lastSearch.Add(1 * time.Hour)
	if today.Before(comparator) {
		return false
	} else {
		return true
	}
}

// Get logo and title of a url
func getPageInfo(url string) (string, string) {
	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		log.Println(err)
		return "", ""
	}
	return s.Preview.Title, s.Preview.Icon
}
