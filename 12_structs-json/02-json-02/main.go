package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type gResult struct {
	Data  string 'json:"name"'
	Valor string 'json:"craft"'
}

func main() {
	uri := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.7418/dados?formato=json"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Erro", err)
		return
	}
	defer resp.Body.Close()
	var gr gResult
	err = json.NewDecoder(resp.Body).Decode(&gr)
	if err != nil {
		log.Println("Erro", err)
		return
	}
	fmt.Println(gr)
}
