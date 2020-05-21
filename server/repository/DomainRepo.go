package repository

import (
	"log"
	"sort"
	"strings"

	"github.com/likexian/whois-go"
)

// const SSL_DEFAULT = "-"
// const LINE_OWNER = 41
// const LINE_COUNTRY = 47

type DomainAPI struct {
	Host      string      `json:"host"`
	Protocol  string      `json:"protocol"`
	Endpoints []ServerAPI `json:"endpoints"`
	Erros     []ErrorsAPI `json:"errors"`
}

type ServerAPI struct {
	IPAddress string `json:"ipAddress"`
	Grade     string `json:"grade"`
}

type ErrorsAPI struct {
	Message string `json:"message"`
}

//https://github.com/ssllabs/research/wiki/SSL-Server-Rating-Guide
var sslGrades = []string{"A", "A+", "B", "C", "D", "E", "F"}

// Return the owner and country using whois(ip)
func (d *DomainAPI) WhoisServerAttributes() (string, string) {
	result, err := whois.Whois(d.Host)
	if err != nil {
		log.Fatal(err)
	}
	owner, country := splitWhois(result)
	return owner, country
}

// Find owner and country in the text that return whois(ip)
func splitWhois(response string) (string, string) {

	var owner, country string
	dataOwner := (strings.Split(response, "Name:"))
	dataCountry := (strings.Split(response, "Country:"))

	if len(dataOwner) > 1 {
		dataOwner = (strings.Split(dataOwner[1], "\n"))
		owner = strings.Trim(dataOwner[0], " ")
	}

	if len(dataCountry) > 1 {
		dataCountry = strings.Split(dataCountry[1], "\n")
		country = strings.Trim(dataCountry[0], " ")
	}

	return owner, country
}

// Seach minor grade in all servers
func (d *DomainAPI) SearchMinorGrade() string {
	servers := d.Endpoints

	// Is there are no servers
	if len(servers) < 1 {
		// return SSL_DEFAULT
		return ""
	}

	minor := servers[0].Grade

	for _, server := range servers {
		// 	grades := strings.Split(servers.SslGrade, "")

		// if len(servers) <= 0 {
		// 	return SSL_DEFAULT
		// }
		grade := server.Grade

		// indice Temporal
		iT := sort.SearchStrings(sslGrades, grade)
		jT := sort.SearchStrings(sslGrades, minor)

		if iT < len(sslGrades) && sslGrades[iT] == grade {
			// if grade != SSL_DEFAULT {

			// Compare indexes that are ordered
			if jT >= iT {
				minor = grade
			}

			// }
		}
	}

	return minor
}
