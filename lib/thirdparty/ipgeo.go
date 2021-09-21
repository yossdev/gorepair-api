package thirdparty

import (
	"encoding/json"
	"gorepair-rest-api/models"
	"io/ioutil"
	"log"
	"net/http"
)

func GetIpLocation() models.IpLocation {
	url := "https://freegeoip.app/json/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var payload models.IpLocation
	err = json.Unmarshal(body, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
	return payload
}