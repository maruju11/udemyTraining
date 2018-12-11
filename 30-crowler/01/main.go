package main

import (
	"fmt"
)

// Fetcher returns the body of URL and
// a slice of URLs found on that page.
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type result struct {
	url, body string
	urls      []string
	err       error
	depth     int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	results := make(chan *result)
	fetched := make(map[string]bool)
	fetch := func(url string, depth int) {
		body, urls, err := fetcher.Fetch(url)
		results <- &result{url, body, urls, err, depth}
	}

	go fetch(url, depth)
	fetched[url] = true

	// 1 url is currently being fetched in background, loop while fetching
	for fetching := 1; fetching > 0; fetching-- {
		res := <-results

		// skip failed fetches
		if res.err != nil {
			fmt.Println(res.err)
			continue
		}

		fmt.Printf("found: %s %q\n", res.url, res.body)

		// follow links if depth has not been exhausted
		if res.depth > 0 {
			for _, u := range res.urls {
				// don't attempt to re-fetch known url, decrement depth
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1)
					fetched[u] = true
				}
			}
		}
	}

	close(results)
}

func main() {
	Crawl("http://fairon.com.br/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := (*f)[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	"http://fairon.com.br/": &fakeResult{
		"785937896",
		[]string{
			"http://fairon.com.br/",
			"http://fairon.com.br/",
		},
	},
	"http://fairon.com.br/contato": &fakeResult{
		"Packages",
		[]string{
			"http://fairon.com.br/",
			"http://fairon.com.br/contato/",
			"http://fairon.com.br/contato/fmt/",
			"http://fairon.com.br/contato/os/",
		},
	},
	"http://fairon.com.br/erp-multiempresa/": &fakeResult{
		"Multiempresa",
		[]string{
			"http://fairon.com.br/",
			"http://fairon.com.br/erp-multiempresa/",
		},
	},
	"http://fairon.com.br/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://fairon.com.br/",
			"http://fairon.com.br/pkg/",
		},
	},
}
