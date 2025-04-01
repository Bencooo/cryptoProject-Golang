package main

import (
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

func getResponse() string {
	resp, err := http.Get("https://api.kraken.com/0/public/Time")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if body != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Body :", body)
	return string(body)
}

func main() {
	getResponse()
}
