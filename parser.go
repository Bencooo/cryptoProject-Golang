package main

import (
	"encoding/json"
	"log"
)

func parseResponse(body string) Result {
	var value Result
	err := json.Unmarshal([]byte(body), &value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func parseAssetPairResponse(body string) AssetPairResponse {
	var value AssetPairResponse
	err := json.Unmarshal([]byte(body), &value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func parseTickerPairResponse(body string) TickerResponse {
	var value TickerResponse
	err := json.Unmarshal([]byte(body), &value)
	if err != nil {
		log.Fatal(err)
	}
	return value
}
