package engine

import (
	"log"
	"test/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetcher %s", r.Url)
	//获取这个请求的Body
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetcher url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	//拿到Body，从中解析出Requests和Items
	return r.Parser.Parse(body, r.Url), nil

}
