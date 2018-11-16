package engine

import (
	"log"
	"test/crawler/fetcher"
)

type SimpleEngine struct{}

func(e SimpleEngine)Run(seeds ...Request) {
	//声明一个任务队列
	var requests []Request

	for _, r := range seeds{
		requests = append(requests,r)
	}
	//循环任务队列
	for len(requests) > 0 {
		//取出一个"请求"，减去这个请求
		r := requests[0]
		requests = requests[1:]

		ParseResult, err := worker(r)

		if err != nil{
			continue
		}
		//在把Request添加到任务队列中
		requests = append(requests, ParseResult.Requests...)
		//打印Parser获取到的Items
		for _, item :=range ParseResult.Items{
			log.Printf("Got Item %v",item)
		}
	}
}

func  worker(r Request)(ParseResult, error){
	//log.Printf("Fetcher %s", r.Url)
	//获取这个请求的Body
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetcher url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	//拿到Body，从中解析出Requests和Items
	return r.ParserFunc(body), nil

}