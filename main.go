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

func getResponse() string {
	resp, err := http.Get("https://api.kraken.com/0/public/Time")
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

func main() {
	response := getResponse()
	result := parseResponse(response)

	fmt.Println(result.Result)
	fmt.Println(result.Result.Unixtime)
	fmt.Println(result.Result.Rfc1123)
}
