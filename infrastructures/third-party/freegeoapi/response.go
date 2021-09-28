package freegeoapi

import _ipgeo "gorepair-rest-api/src/ip-geo"

type Response struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	RegionCode  string  `json:"region_code"`
	RegionName  string  `json:"region_name"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	MetroCode   int     `json:"metro_code"`
}

func (resp *Response) ToDomain() _ipgeo.Domain {
	return _ipgeo.Domain{
		IP: 			resp.IP,
		City: 			resp.City,
		CountryName: 	resp.CountryName,
		RegionName: 	resp.RegionName,
	}
}