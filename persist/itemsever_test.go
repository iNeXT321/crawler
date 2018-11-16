package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"test/crawler/engine"
	"test/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {

	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1275821982",
		Type:"zhenai",
		Id:"1275821982",
		PayLoad: model.Profile{
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
	//保存一个Item
	err := save(expected)
	if err != nil{
		panic(err)
	}
	//拿保存的Item
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().Index("dating_profile").Type(expected.Type).Id(expected.Id).Do(context.Background())
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s", *resp.Source)
	//验证Item
	var actual engine.Item
	//解编成想要的结构体
	json.Unmarshal(*resp.Source, &actual)
	actualProfile,_ := model.FromJsonObj(actual.PayLoad)
	actual.PayLoad = actualProfile

	if actual != expected{
		t.Errorf("got %v; expected %v ", actual, expected)
	}
}
