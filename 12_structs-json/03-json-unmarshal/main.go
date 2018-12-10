package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Artist comments ... jçadjf
type Artist struct {
	ID             int
	Name           string
	ResourceURL    string
	ReleasesURL    string
	URI            string
	Realname       string
	Profile        string
	DataQuality    string
	Namevariations []string
	Aliases        []struct {
		ID          int
		Name        string
		ResourceURL string
	}
	Urls   []string
	Images []struct {
		Type        string
		Width       int
		Height      int
		URI         string
		URI150      string
		ResourceURL string
	}
}

func main() {
	res, _ := http.Get("https://api.discogs.com/artists/1373")

	temp, _ := ioutil.ReadAll(res.Body)
	var artist Artist
	err := json.Unmarshal(temp, &artist)
	if err != nil {
		fmt.Println("Existe um erro", err)
	}
	fmt.Println(artist.Name)
	fmt.Println(artist.Images[1].Height)
	fmt.Println(artist)
}
