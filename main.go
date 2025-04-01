package main

type Result struct {
	Error  []string
	Result struct {
		Unixtime int    `json:"unixtime"`
		Rfc1123  string `json:"rfc1123"`
	} `json:"result"`
}

func main() {

}
