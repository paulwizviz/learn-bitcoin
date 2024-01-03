package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type rpcbody struct {
	JsonRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      string `json:"id"`
}

func main() {

	body := rpcbody{
		JsonRPC: "1.0",
		Method:  "getblockchaininfo",
		ID:      "go-client",
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal("Marshal JSON", err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:18443/", bytes.NewReader(b))
	if err != nil {
		log.Fatal("New request", err)
	}
	req.SetBasicAuth("username", "password")
	req.Header.Set("Content-Type", "application/json")

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal("Execute request", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading body", err)
	}
	fmt.Printf("Body: %s", data)
}
