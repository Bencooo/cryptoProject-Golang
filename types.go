package main

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
