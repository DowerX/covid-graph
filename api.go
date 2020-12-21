package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Event struct {
	//Country     string `json:"Country"`
	//CountryCode string `json:"CountryCode"`
	//Lat         string `json:"Lat"`
	//Lon         string `json:"Lon"`
	Cases int `json:"Cases"`
	//Status      string `json:"Status"`
	Date string `json:"Date"`
}

func getDayOne(country string, status string) ([]Event, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	resp, err := http.Get(fmt.Sprintf("https://api.covid19api.com/dayone/country/%s/status/%s", country, status))
	if err != nil {
		return nil, err
	}
	var arr []Event
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}
