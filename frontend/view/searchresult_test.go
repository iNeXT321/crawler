package view

import (
	"os"
	"test/crawler/engine"
	"test/crawler/frontend/model"
	common "test/crawler/model"
	"testing"
)

func TestSearchResultView(t *testing.T)  {
	view := CreateSearchResultView("template.html")
	page := model.SearchResult{}
	page.Hits = 123
	item :=  engine.Item{
		Url:"http://album.zhenai.com/u/1275821982",
		Type:"zhenai",
		Id:"1275821982",
		PayLoad: common.Profile{
			Name:       "雪儿",
			Gender:     "女",
			Marriage:   "是",
			Age:        "20",
			XinZuo:     "魔羯座(12.22-01.19)",
			Height:     "170",
			Weight:     "60",
			WorkPlace:  "杭州",
			Income:     "6000",
			Occupation: "IT",
			Education:  "本科",

			HoKou: "杭州",
			House: "有",
			Car:   "有",
		},
	}

	for i := 0; i < 100; i++{
		page.Items = append(page.Items, item)
	}
	out, err := os.Create("template.test.html")
	err = view.Render(out, page)
	if err != nil{
		panic(err)
	}
}
