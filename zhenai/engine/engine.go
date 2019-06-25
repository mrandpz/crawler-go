package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, e := worker(r)
		if e != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}

	}
}

// type Request struct {
//	Url        string
//	ParserFunc func([]byte) ParseResult
//}

//type ParseResult struct {
//	Requests []Request
//	Items    []interface{}
//}
func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error"+"Fetchering url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
