package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jehm09/Android-Queries/server/model"
)

//This API gives us information about the domain, its servers and the ssl grade of each server
const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

type ActualData struct {
	History *model.History
	Domain  *model.Domain
}

func NewActualData(history *model.History, domain *model.Domain) *ActualData {
	return &ActualData{History: history, Domain: domain}
}

func GetDomain(host string, actualData *ActualData) *model.Domain {
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

	return &domainA
}

func createDomain(domainA DomainAPI, actualData *ActualData) *models.Domain {
	// result := &models.Domain{}
	if actualData.History != nil {

	}

}
