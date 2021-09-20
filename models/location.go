package models

type RevOutput struct {
	Items []struct {
		Title string `json:"title"`
	} `json:"items"`
}