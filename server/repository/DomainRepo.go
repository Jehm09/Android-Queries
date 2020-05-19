package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//This API gives us information about the domain, its servers and the ssl grade of each server
const API_DOMAINS_URL = "https://api.ssllabs.com/api/v3/analyze?host="

func GetDomain(host string) *DomainAPI {
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
