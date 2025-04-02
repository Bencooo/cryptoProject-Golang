package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
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
	// Remove DB file if it exists
	err := os.Remove("kraken_database.db")
	if err != nil {
		log.Fatal(err)
	}

	// Open a new database connection
	db, err := sql.Open("sqlite3", "kraken_database.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// Create table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS system_status (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		unixtime INTEGER,
		rfc1123 TEXT
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("System status table created successfully")
	log.Println("-------------------------")

	//create a table for asset pairs
	createAssetTableSQL := `
	CREATE TABLE IF NOT EXISTS asset_pairs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pair TEXT NOT NULL UNIQUE,
		altname TEXT,
		base TEXT,
		quote TEXT
	);`

	_, err = db.Exec(createAssetTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Asset pairs table created successfully")
	log.Println("-------------------------")

	// Create table for ticker info
	createTickerTableSQL := `
	CREATE TABLE IF NOT EXISTS ticker_info (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		pair TEXT NOT NULL,
		ask_price TEXT,
		bid_price TEXT,
		last_trade_price TEXT,
		open_price TEXT,
		high_24h TEXT,
		low_24h TEXT,
		volume_24h TEXT,
		retrieved_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(pair) REFERENCES asset_pairs(pair)
	);`

	_, err = db.Exec(createTickerTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Ticker info table created successfully")
	log.Println("-------------------------")

	// Get the response from the API
	body := getResponse("https://api.kraken.com/0/public/Time")
	// Parse the response
	systemStatusResponse := parseResponse(body)
	// Insert the data into the database
	insertSQL := `INSERT INTO system_status (unixtime, rfc1123) VALUES (?, ?)`
	_, err = db.Exec(insertSQL, systemStatusResponse.Result.Unixtime, systemStatusResponse.Result.Rfc1123)
	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}
	log.Println("System status data inserted successfully")
	log.Println("-------------------------")

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
