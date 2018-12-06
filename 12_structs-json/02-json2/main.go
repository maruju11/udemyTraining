package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Gresult recebe os dados seriais do governo e retorna data e valor em milhoes de reais
type Gresult []struct {
	Data  string
	Valor string
}

func main() {
	uri := "https://api.bcb.gov.br/dados/serie/bcdata.sgs.7418/dados?formato=json"

	res, err := http.Get(uri)
	if err != nil {
		log.Println("Erro", err)
		return
	}
	defer res.Body.Close()

	//temp, _ := ioutil.ReadAll(res.Body)

	var gr Gresult
	//err = json.Unmarshal(res.Body, &gr)
	err = json.NewDecoder(res.Body).Decode(&gr)
	if err != nil {
		log.Println("Erro", err)
		return
	}
	fmt.Println(gr)
	fmt.Println(gr[1].Data)
}
