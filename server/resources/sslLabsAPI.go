package resources

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"truora-server/models"
)

const url = "https://api.ssllabs.com/api/v3/analyze?host="

func GetInfoDomainFromAPI(domain string) models.SSLLabs {
	fullUrl := url + domain

	spaceCliente := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequest(http.MethodGet, fullUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "developer-test")
	res, err := spaceCliente.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	decoder := json.NewDecoder(res.Body)
	var ssllab models.SSLLabs
	err = decoder.Decode(&ssllab)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	return ssllab
}
