package parser

import (
	"regexp"
	"test/crawler/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
const genderRe = `<span class="grayL">性别：</span>([^<]+)</td>`
const cityUrlRe = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`

func ParseCity(contents []byte) engine.ParseResult {
	cityRegular := regexp.MustCompile(cityRe)
	cityMatches := cityRegular.FindAllSubmatch(contents, -1)

	genderRegular := regexp.MustCompile(genderRe)
	genderMatches := genderRegular.FindAllSubmatch(contents, -1)

	cityUrlRegular := regexp.MustCompile(cityUrlRe)
	cityUrlMatches := cityUrlRegular.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for i := 0; i < len(cityMatches); i++ {
		userName := string(cityMatches[i][2])
		userGender := string(genderMatches[i][1])
		url :=  string(cityMatches[i][1])
		//result.Items = append(result.Items, "User "+userName+" Gender "+userGender)
		result.Requests = append(result.Requests, engine.Request{
			Url: url,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, url, userName, userGender)
			},
		})
	}

	for _, m := range cityUrlMatches{
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}

	return result
}
