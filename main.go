package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

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
		fmt.Printf("\n• %s | Altname: %s | Base: %s | Quote: %s\n", pairName, details.Altname, details.Base, details.Quote)

		url := fmt.Sprintf("https://api.kraken.com/0/public/Ticker?pair=%s", pairName)
		tickerResponse := getResponse(url)
		tickerParse := parseTickerPairResponse(tickerResponse)

		if tickerData, ok := tickerParse.Result[pairName]; ok {
			fmt.Printf("  ➜ Ask: %s | Bid: %s | Last: %s | Open: %s\n",
				tickerData.A[0], tickerData.B[0], tickerData.C[0], tickerData.O)
		} else {
			fmt.Println("Unknown")
		}

		count++
		if count == 10 {
			break
		}
	}

}
