package controller

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"test/crawler/engine"
	"test/crawler/frontend/model"
	"test/crawler/frontend/view"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//初始化SearchResultHandler
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

//localhost:8888/search/?q=男 已购房 &from=20
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//去空格
	q := strings.TrimSpace(req.FormValue("q"))
	//字符串转成int
	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = h.view.Render(w, page)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

}
//向数据库查询数据,并向 result中装载数据
func (h SearchResultHandler) getSearchResult(q string, from int)(model.SearchResult,error){
	var result model.SearchResult
	result.Query = q
	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil{
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	//分页
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string{
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "PayLoad.$1:")
}