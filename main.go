package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Result struct {
	Error  []string
	Result struct {
		Unixtime int    `json:"unixtime"`
		Rfc1123  string `json:"rfc1123"`
	} `json:"result"`
}

type AssetPairResponse struct {
	Error  []string                `json:"error"`
	Result map[string]AssetDetails `json:"result"`
}

type AssetDetails struct {
	Altname string `json:"altname"`
	Base    string `json:"base"`
	Quote   string `json:"quote"`
}

type TickerResponse struct {
	Error  []string                `json:"error"`
	Result map[string]TickerByPair `json:"result"`
}

type TickerByPair struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []int    `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}

func getResponse(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Body :", string(body))
	return string(body)
}

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

func main() {
	response := getResponse("https://api.kraken.com/0/public/Time")
	result := parseResponse(response)

	fmt.Println(result.Result)
	fmt.Println(result.Result.Unixtime)
	fmt.Println(result.Result.Rfc1123)

	response2 := getResponse("https://api.kraken.com/0/public/AssetPairs")
	resultAssetPair := parseAssetPairResponse(response2)

	count := 0
	for pairName, details := range resultAssetPair.Result {
		fmt.Printf("• %s | Altname: %s | Base: %s | Quote: %s\n", pairName, details.Altname, details.Base, details.Quote)
		count++
		if count == 10 {
			return
		}
	}
}
