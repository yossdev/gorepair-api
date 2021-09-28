package realemailapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RealEmailResponse struct {
	Status string `json:"status"`
}

func RealEmail(email string) *RealEmailResponse {
	url := "https://isitarealemail.com/api/email/validate?email=" + url.QueryEscape(email)
	req, _ := http.NewRequest("GET", url, nil)
	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil
	}

	var resp RealEmailResponse
	json.Unmarshal(body, &resp)

	fmt.Printf("status for %v is %v", email, resp.Status)
	return &resp
}