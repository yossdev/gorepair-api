package freegeoapi

import (
	"encoding/json"
	_ipgeo "gorepair-rest-api/src/ip-geo"
	"io/ioutil"
	"log"
	"net/http"
)

type IpAPI struct {
	httpClient http.Client
}

func NewIpAPI() _ipgeo.Repository {
	return &IpAPI{
		httpClient: http.Client{},
	}
}

func (ipl *IpAPI) GetLocationByIP() (_ipgeo.Domain, error) {
	url := "https://freegeoip.app/json/"
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return _ipgeo.Domain{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	payload := Response{}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
		return _ipgeo.Domain{}, err
	}

	return payload.ToDomain(), err
}