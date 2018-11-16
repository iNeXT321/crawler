package parser

import (
	"regexp"
	"test/crawler/engine"
	"test/crawler/model"
)

const Re = `<div class="m-btn purple" [^>]*>([^<]+)</div>`
const IdRe = `<div class="id" [^>]*>ID：([^<]+)</div>`
func ParseProfile(contents []byte,url string, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	re := regexp.MustCompile(Re)
	matches := re.FindAllSubmatch(contents, -1)

	var result engine.ParseResult
	//条件不满足9个则舍弃
	if len(matches) != 9{
		return result
	}

	idRe := regexp.MustCompile(IdRe)
	IdMatches := idRe.FindAllSubmatch(contents,-1)


	var box [9]string
	var userId string
	for i, m := range matches {
		box[i] = string(m[1])
	}
	for _, m := range IdMatches{
		userId = string(m[1])
	}

	profile.Marriage = box[0]
	profile.Age = box[1]
	profile.XinZuo = box[2]
	profile.Height = box[3]
	profile.Weight = box[4]
	profile.WorkPlace = box[5]
	profile.Income = box[6]
	profile.Occupation = box[7]
	profile.Education = box[8]

	result = engine.ParseResult{
		Items:[]engine.Item{
			{
				Url:url,
				Type:"zhenai",
				Id:userId,
				PayLoad:profile,
			},
		},
	}

	return result
}
