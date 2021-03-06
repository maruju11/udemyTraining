package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.amazon.com")
	if err != nil {
		log.Fatalln(err)
	}
	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "lendo entrada", err)
	}
	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 10 {
			break
		}
		i++
	}
}
